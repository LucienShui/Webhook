/*
@File: history.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-25 01:33  Lucien     1.0         Init
*/
package model

import (
	"errors"
	"time"
)

type History struct {
	Name      string `gorm:"primary_key;type:varchar(32)"`
	Content   string `gorm:"type:mediumtext"`
	UpdatedAt time.Time
}

func (object History) Save() error {
	if object.Name == "" {
		return errors.New("empty name")
	}
	if db.NewRecord(&object) {
		return db.Create(&object).Error
	}
	return db.Save(&object).Error
}

func (object *History) Load() error {
	return db.Find(&object, "`name` = ?", object.Name).Error
}
