#!/usr/bin/env bash
./install.sh

if [ $? -ne 0 ]; then
    exit 0
fi

gitclone https://github.com/WindomZ/gitclone.git
