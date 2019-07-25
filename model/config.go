/*
@File: config.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-25 01:33  Lucien     1.0         Init
*/
package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
)

type Config struct {
	Address string `json:"address"`
	Port uint16 `json:"port"`
	Webhooks []Webhook `json:"webhooks"`
}

func (config *Config) Load(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Fatal(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Fatal(err)
	}
	configPrinter(*config)
}

func configPrinter(config Config) {
	buf := bytes.NewBufferString("config loaded:\n")
	buf.WriteString(fmt.Sprintf("Address: %s\n", config.Address))
	buf.WriteString(fmt.Sprintf("Port: %d\n", config.Port))
	buf.WriteString("Records: [")
	length := len(config.Webhooks)
	for i := 0; i < length - 1; i++ {
		buf.WriteString(fmt.Sprintf("%s, ", config.Webhooks[i].Name))
	}
	buf.WriteString(fmt.Sprintf("%s]\n", config.Webhooks[length - 1].Name))
	logger.Info(buf.String())
}

func (config *Config) Get(name string) (Webhook, error) {
	for _, webhook := range config.Webhooks {
		if webhook.Name == name {
			return webhook, nil
		}
	}
	return Webhook{}, errors.New("not exist")
}
