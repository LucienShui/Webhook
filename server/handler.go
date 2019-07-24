/*
@File: handler.go
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
	"net/http"
	"os/exec"
)

func cmd(script string) {
	cmd := exec.Command("bash", "-x", script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err) // TODO
	} else {
		fmt.Println(string(out))
	}
}

func request(requests *gin.Context, config model.Config) {
	name := requests.Param("name")
	password := requests.DefaultQuery("keyword", "")
	webhook, err := config.Get(name)
	if err != nil {
		panic(err) // TODO
	}
	if password == webhook.Password {
		go cmd(webhook.Script)
	}
	requests.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"name":    webhook.Name,
	})
}
