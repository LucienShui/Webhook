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
	cmd := exec.Command("bash", "-c", webhook.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Info("execute command: '%s' failed", webhook.Script)
		logger.Warn(err)
	} else {
		logger.Info("command [%s] output:\n%s", webhook.Script, string(out))
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
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusAccepted,
				"message": "success",
			})
			go cmd(webhook)
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "password wrong",
			})
			logger.Warn("password wrong")
		}
	}
}

func queryLog(context *gin.Context) {
	history := model.History{Name: context.Param("name")}
	if err := history.Load(); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err,
		})
		logger.Warn(err)
	} else {
		rawRequest := context.DefaultQuery("raw", "false")
		if rawRequest == "true" {
			context.String(http.StatusOK,
				"[%s]\n%s", history.UpdatedAt.Format("2006-01-02 15:04:05"), history.Content)
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"message": history.Content,
				"time": history.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}
}

func notFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
	})
}
