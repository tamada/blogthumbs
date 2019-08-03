package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"

	flag "github.com/spf13/pflag"
)

const VERSION = "1.0.0"

type Formatter interface {
	format(fileName, thumbnail string) string
}

func getVersionString() string {
	return fmt.Sprintf("blogthumbs version %s", VERSION)
}

func getUsage() string {
	return fmt.Sprintf(`Usage: blogthumbs [OPTIONS] <IMAGES...>
OPTIONS
    -d, --dest [DIR]     specifies the destination.
    -s, --size [SIZE]    specifies the size of resultant image. Default is 240.
    -t, --type [TYPE]    specifies the type of blog system.
                         available values: 'markdown', and 'hugo'. default is markdown.
    -h, --help           print this message.
    -v, --version        print version of blogthums.
IMAGES
    specifies image for resizing. Acceptable format is 'gif', 'jpeg', and 'png'.`)
}

type options struct {
	dest        string
	system      string
	size        int
	helpFlag    bool
	versionFlag bool
	args        []string
}

func buildFlagSet(args []string) (*flag.FlagSet, *options) {
	var flags = flag.NewFlagSet("blogthumbs", flag.ContinueOnError)
	var opts = options{}
	flags.Usage = func() { fmt.Println(getUsage()) }
	flags.StringVarP(&opts.dest, "dest", "d", ".", "specifies the destination")
	flags.StringVarP(&opts.system, "type", "t", "markdown", "specifies the type of blog system")
	flags.IntVarP(&opts.size, "size", "s", 240, "specifies the size of resultant image.")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this messgae")
	flags.BoolVarP(&opts.versionFlag, "version", "v", false, "print blogthumbs version")
	return flags, &opts
}

func parseArgs(args []string) (*options, error) {
	var flag, opts = buildFlagSet(args)
	if err := flag.Parse(args); err != nil {
		return nil, err
	}
	opts.args = flag.Args()
	return opts, nil
}

func buildFormatter(system string) Formatter {
	system = strings.ToLower(system)
	if system == "hugo" {
		return &Hugo{}
	}
	return &Markdown{}
}

func findDestination(source string, opts *options) string {
	var dir, file = filepath.Split(source)
	if opts.dest == "." {
		return fmt.Sprintf("%s/t_%s", dir, file)
	}
	var dest = strings.TrimRight(opts.dest, "/")
	return fmt.Sprintf("%s", filepath.Join(dest, file))
}

func findFormat(source string) (string, error) {
	var file, err = os.Open(source)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var _, format, err2 = image.DecodeConfig(file)
	return format, err2
}

func readImage(source string) (image.Image, string, error) {
	var format, err = findFormat(source)
	if err != nil {
		return nil, "", err
	}
	var file, err2 = os.Open(source)
	if err2 != nil {
		return nil, "", err
	}
	defer file.Close()
	switch format {
	case "jpeg":
		var img, err = jpeg.Decode(file)
		return img, format, err
	case "png":
		var img, err = png.Decode(file)
		return img, format, err
	case "gif":
		var img, err = gif.Decode(file)
		return img, format, err
	}
	return nil, "", fmt.Errorf("%s: unknown image format", source)
}

func scaledSize(image image.Image, imageSize int) (int, int) {
	var dim = image.Bounds()
	var size = dim.Dx()
	if size > dim.Dy() {
		size = dim.Dy()
	}
	var scale = float32(imageSize) / float32(size)
	return int(float32(dim.Dx()) * scale), int(float32(dim.Dy()) * scale)
}

func writeImage(destImage image.Image, format string, destPath string) (string, error) {
	var dest, err = os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()
	switch format {
	case "jpeg":
		if err := jpeg.Encode(dest, destImage, &jpeg.Options{Quality: 100}); err != nil {
			return "", err
		}
	case "png":
		if err := png.Encode(dest, destImage); err != nil {
			return "", err
		}
	case "gif":
		if err := gif.Encode(dest, destImage, nil); err != nil {
			return "", err
		}
	}
	return destPath, nil
}

func createThumbnail(source string, opts *options) (string, error) {
	var destFile = findDestination(source, opts)
	var sourceImage, format, err = readImage(source)
	if err != nil {
		return "", err
	}
	var width, height = scaledSize(sourceImage, opts.size)
	var destImage = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(destImage, destImage.Bounds(), sourceImage, sourceImage.Bounds(), draw.Over, nil)
	return writeImage(destImage, format, destFile)
}

func storeThumbnail(source string, opts *options, formatter Formatter) (string, error) {
	var thumbnail, err = createThumbnail(source, opts)
	if err != nil {
		return "", err
	}
	return formatter.format(source, thumbnail), nil
}

func perform(opts *options) int {
	var formatter = buildFormatter(opts.system)
	for _, arg := range opts.args[1:] {
		var output, err = storeThumbnail(arg, opts, formatter)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Println(output)
	}
	return 0
}

func goMain(args []string) int {
	var opts, err = parseArgs(args)
	if err != nil {
		fmt.Printf(err.Error())
		return 1
	}
	if opts.helpFlag {
		fmt.Println(getUsage())
	}
	if opts.versionFlag {
		fmt.Println(getVersionString())
	}
	if opts.helpFlag || opts.versionFlag {
		return 0
	}
	return perform(opts)
}

func main() {
	var status = goMain(os.Args)
	os.Exit(status)
}
