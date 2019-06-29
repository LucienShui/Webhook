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

type Job struct {
	Name      string `gorm:"type:varchar(32);primary key;index;" json:"name"`
	Script    string `gorm:"type:text" json:"script"`
	Password  string `gorm:"type:varchar(16)" json:"password"`
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (object *Job) Get() error {
	return db.Find(&object, "`name` = ?", object.Name).Error
}

func (object *Job) Save() error {
	return db.Create(&object).Error
}

func (object *Job) Add(content string) error {
	history := History{
		Name:    object.Name,
		Content: content,
	}
	return history.Save()
}
