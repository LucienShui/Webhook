/*
@File: data.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-25 23:40  Lucien     1.0         Init
*/
package model

import (
	"time"
)

type Data struct {
	Name string `gorm:"primary_key"`
	Script string `gorm:"type:text"`
	Password string `gorm:"type:varchar(16)"`
	History []History
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (object *Data) Get() error {
	return db.Find(&object, "`name` = ?", object.Name).Error
}

func (object *Data) Save() error {
	return db.Create(&object).Error
}


