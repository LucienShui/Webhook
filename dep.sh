#!/usr/bin/env bash

DEP="
github.com/gin-gonic/gin
github.com/jinzhu/gorm
github.com/mattn/go-sqlite3
github.com/wonderivan/logger
"

for each in ${DEP}; do
    /usr/bin/env go get -u ${each}
done
