[![Build Status](https://travis-ci.com/tamada/blogthumbs.svg?branch=master)](https://travis-ci.com/tamada/blogthumbs)
[![Coverage Status](https://coveralls.io/repos/github/tamada/blogthumbs/badge.svg?branch=master)](https://coveralls.io/github/tamada/blogthumbs?branch=master)
[![codebeat badge](https://codebeat.co/badges/9aea5795-9f10-4dc2-b63b-d4e12f3aed3f)](https://codebeat.co/projects/github-com-tamada-blogthumbs-master)
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
