/*
@File: model.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-25 01:39  Lucien     1.0         Init
*/
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "webhook.db")
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&History{}) {
		if err := db.CreateTable(&History{}).Error; err != nil {
			panic(err)
		}
	}
}
