# gitclone

A cli tool, git clone repository in the `go get` style.

## Features

- [x] git clone repository in the `go get` style
- [ ] command: clone - same as `git clone`
- [ ] command: list - prints a list of files in the current directory
- [ ] command: mkdir - create a default directory
- [ ] command: search - search repositories witch in current directory
- [ ] command: import - `gitclone` a repository in the directory to current directory

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

So easy to use!

Result:
```
|- .
|- ..
|- github.com
|  - WindomZ
|     - gitclone
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
