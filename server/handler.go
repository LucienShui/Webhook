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

func cmd(webhook model.Webhook) {
	cmd := exec.Command("bash", "-c", webhook.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err) // TODO
	} else {
		fmt.Println(string(out))
		history := model.History{
			Name: webhook.Name,
			Content: string(out),
		}
		if err := history.Save(); err != nil {
			panic(err) // TODO
		}
	}
}

func request(context *gin.Context) {
	name := context.Param("name")
	password := context.DefaultQuery("password", "")
	webhook, err := config.Get(name)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": "record not found",
		})
	} else {
		if password == webhook.Password {
			go cmd(webhook)
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusAccepted,
			"message": "success",
		})
	}
}

func notFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
	})
}
