/*
@File: model.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-07-24 23:29  Lucien     1.0         Init
*/
package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Webhook struct {
	Name string `json:"name"`
	Script string `json:"script"`
	Password string `json:"password"`
}

type Config struct {
	Address string `json:"address"`
	Port uint16 `json:"port"`
	Webhooks []Webhook `json:"webhooks"`
}

func (config *Config) Load(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) // TODO
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err) // TODO
	}
}

func (config *Config) Get(name string) (Webhook, error) {
	for _, webhook := range config.Webhooks {
		if webhook.Name == name {
			return webhook, nil
		}
	}
	return Webhook{}, errors.New("not exist")
}
