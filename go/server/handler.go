/*
@File: handler.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-25 22:29  Lucien     1.0         Init
*/
package server

import (
	"github.com/LucienShui/Webhook/go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func cmd(object model.Job) {
	cmd := exec.Command("bash", "-x", object.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err) // TODO
	} else {
		if err := object.Add(string(out)); err != nil {
			panic(err) // TODO
		}
	}
}

func request(requests *gin.Context) {
	object := model.Job{Name: requests.Param("name")}
	if err := object.Get(); err != nil {
		panic(err) // TODO
	} else {
		go cmd(object)
		requests.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"name":    object.Name,
		})
	}
}

func create(requests *gin.Context) {
	job := model.Job{Name: requests.Param("name")}
	if err := requests.ShouldBindJSON(&job); err != nil {
		panic(err) // TODO
	} else {
		if err := job.Save(); err != nil {
			panic(err) // TODO
		} else {
			requests.JSON(http.StatusCreated, gin.H{
				"code":    http.StatusCreated,
				"message": "create success",
			})
		}
	}
}
