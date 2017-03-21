#!/usr/bin/env bash
go install

if [ $? -ne 0 ]; then
    exit 0
fi

#gitclone https://github.com/WindomZ/gitclone.git
gitclone help
