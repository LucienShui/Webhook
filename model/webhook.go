/*
@File: model.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-24 23:29  Lucien     1.0         Init
*/
package model

type Webhook struct {
	Name string ` json:"name"`
	Script string `json:"script"`
	Password string `json:"password"`
}

