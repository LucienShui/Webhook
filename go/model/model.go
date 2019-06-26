/*
@File: model.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-25 23:28  Lucien     1.0         Init
*/
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic(err)
	}
}

