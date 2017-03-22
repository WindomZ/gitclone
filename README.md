# gitclone
[![Build Status](https://travis-ci.org/WindomZ/gitclone.svg?branch=master)](https://travis-ci.org/WindomZ/gitclone)
![License](https://img.shields.io/badge/license-MIT-green.svg)

![v0.5.0](https://img.shields.io/badge/version-v0.5.0-orange.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

A cli tool, git clone repository in the `go get` style.

## Features

- [x] gitclone - git clone repository in the `go get` style
- [x] gitclone list - prints a list of repositories witch in the current directory
- [x] gitclone search - search repositories witch in current directory
- [ ] gitclone link - `gitclone` a repository from the directory to current directory

## Installation

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
    gitclone link
    gitclone <repo>
    gitclone -h | --help
    gitclone --version
```

## Example

`git clone`(or `git pull`) this repo(`github.com/WindomZ/gitclone`), 
simply run:
```bash
gitclone https://github.com/WindomZ/gitclone.git
```

Finally, get the directory structure:
```
|- .
|- ..
|- github.com
|  - WindomZ
|     - gitclone
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
