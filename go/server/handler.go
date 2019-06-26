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
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func cmd(script string) {
	cmd := exec.Command("sh", script)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}
}

func post(requests *gin.Context) {
	go cmd()
	name := requests.Param("name")
	requests.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"name": name,
	})
}