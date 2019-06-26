/*
@File: main.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-25 22:22  Lucien     1.0         Init
*/
package main

import (
	"github.com/LucienShui/Webhook/go/server"
)

func main() {
	if err := server.Run("0.0.0.0", 8000); err != nil {
		panic(err)
	}
}
