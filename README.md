# gitclone
[![Build Status](https://travis-ci.org/WindomZ/gitclone.svg?branch=master)](https://travis-ci.org/WindomZ/gitclone)
![License](https://img.shields.io/badge/license-MIT-green.svg)

![v0.2.0](https://img.shields.io/badge/version-v0.2.0-orange.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

A cli tool, git clone repository in the `go get` style.

## Features

- [x] gitclone - git clone repository in the `go get` style
- [ ] gitclone clone - same as `git clone`
- [ ] gitclone list - prints a list of repositories witch in the current directory
- [ ] gitclone search - search repositories witch in current directory
- [ ] gitclone import - `gitclone` a repository from the directory to current directory

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
NAME:
   gitclone - A cli tool, git clone repository in the `go get` style.

USAGE:
   gitclone [global options] command [command options] [arguments...]

VERSION:
   0.2

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url URL, -u URL, -- URL  git repository URL
   --help, -h                 show help
   --version, -v              print the version
```

## Example

`git clone`(or `git pull`) this repo(`github.com/WindomZ/gitclone`), 
simply run:
```bash
gitclone https://github.com/WindomZ/gitclone.git
```
or
```bash
gitclone -u https://github.com/WindomZ/gitclone.git
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
