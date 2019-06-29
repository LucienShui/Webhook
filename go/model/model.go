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
	"os"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic(err)
	}
	if os.Getenv("GIN_MODE") != "release" {
		db = db.Debug()
	}
	if !db.HasTable(&Job{}) {
		if err := db.CreateTable(&Job{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&History{}) {
		if err := db.CreateTable(&History{}).Error; err != nil {
			panic(err)
		}
	}
}
