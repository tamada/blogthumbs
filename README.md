[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/blogthumbs/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/blogthumbs/releases/tag/v1.0.0)

# blogthumbs

create thumbnail images from given images for blogs.

## Usage

```sh
Usage: blogthumbs [OPTIONS] <IMAGES...>
OPTIONS
    -d, --dest [DIR]     specifies the destination.
    -s, --size [SIZE]    specifies the size of resultant image. Default is 240.
    -t, --type [TYPE]    specifies the type of blog system.
                         available values: 'markdown', and 'hugo'. default is markdown.
    -h, --help           print this message.
    -v, --version        print version of blogthums.
IMAGES
    specifies image for resizing. Acceptable format is 'gif', 'jpeg', and 'png'.
```

## License

[WTFPL](https://github.com/tamada/blogthumbs/blob/master/LICENSE)
