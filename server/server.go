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
	"github.com/LucienShui/Webhook/model"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var config = model.Config{}

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.GET("/:name", request)
	router.NoRoute(notFound)
}

func Run(filename string) error {
	config.Load(filename)
	return router.Run(fmt.Sprintf("%s:%d", config.Address, config.Port))
}
