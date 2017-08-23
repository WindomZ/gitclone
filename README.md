# gitclone

[![Build Status](https://travis-ci.org/WindomZ/gitclone.svg?branch=master)](https://travis-ci.org/WindomZ/gitclone)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gitclone)](https://goreportcard.com/report/github.com/WindomZ/gitclone)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

> A cli tool, git clone repository in the `go get` style.

![v0.7.0](https://img.shields.io/badge/version-v0.7.0-brightgreen.svg)
![status](https://img.shields.io/badge/status-beta-brightgreen.svg)

## Features

- [x] gitclone - git clone repository in the `go get` style
- [x] gitclone list - prints a list of repositories witch in the current directory
- [x] gitclone search - search repositories witch in current directory
- [x] gitclone link - `gitclone` a repository from the directory to current directory

## Install

To install cli, run:
```
$ go get github.com/WindomZ/gitclone
```

Make sure your `PATH` includes the `$GOPATH/bin` directory,
if NOT you can run:
```
export PATH=$PATH:$GOPATH/bin
```

## Usage

```bash
$ gitclone -h

  A cli tool, git clone repository in the `go get` style.

  Usage:
    gitclone (list|ls)
    gitclone search <key>
    gitclone link <filepath>
    gitclone <repo> [--depth=<depth>]
    gitclone -h|--help
    gitclone -v|--version

  Options:
    -h --help     output usage information
    -v --version  output the version number

  Commands:
    list|ls       prints a list of repositories witch in the current directory
    search        search repositories witch in current directory
    link          `gitclone` a repository from the directory to current directory
```

`git clone`(or `git pull`) this repo(`github.com/WindomZ/gitclone`), 
simply run:
```bash
gitclone https://github.com/WindomZ/gitclone.git
```

finally, get the directory structure:
```
|- .
|- ..
|- github.com
|  - WindomZ
|     - gitclone
```

If with `--depth=2`, 
```bash
gitclone https://github.com/WindomZ/gitclone.git --depth=2
```

then get the directory structure:
```
|- .
|- ..
|- WindomZ
|  - gitclone
```

## PowerBy
[go-commander](https://github.com/WindomZ/go-commander) - The solution for building command shell programs

## Related
[WindomZ/gitget](https://github.com/WindomZ/gitget) - Written in Node.js

## Contributing
Welcome to pull requests, report bugs, suggest ideas and discuss **gitclone**, 
i would love to hear what you think about **gitclone** on [issues page](https://github.com/WindomZ/gitclone/issues).

If you like it then you can put a :star: on it.

## License

[MIT](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
