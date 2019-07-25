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
	"github.com/LucienShui/Webhook/model"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"os/exec"
)

func cmd(webhook model.Webhook) {
	logger.Info("execute command: '%s'", webhook.Script)
	cmd := exec.Command("bash", "-c", webhook.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warn(err)
	} else {
		logger.Info("command output:\n%s", string(out))
		history := model.History{
			Name: webhook.Name,
			Content: string(out),
		}
		if err := history.Save(); err != nil {
			logger.Warn(err)
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
