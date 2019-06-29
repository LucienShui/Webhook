/*
@File: history.go
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

type History struct {
	Name      string `gorm:"type:varchar(32)"`
	Content   string `gorm:"type:mediumtext"`
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (object *History) Save() error {
	return db.Create(&object).Error
}
