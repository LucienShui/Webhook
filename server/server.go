/*
@File: server.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-24 23:45  Lucien     1.0         Init
*/
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.GET("/:name", request)
}

func Run(address string, port uint16) error {
	return router.Run(fmt.Sprintf("%s:%d", address, port))
}
