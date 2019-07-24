/*
@File: main.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-24 23:14  Lucien     1.0         Init
*/
package main

import (
	"fmt"
	"github.com/LucienShui/Webhook/model"
)

func main() {
	config := model.Config{}
	config.Load("config.json")
	fmt.Printf("%+v", config)
}
