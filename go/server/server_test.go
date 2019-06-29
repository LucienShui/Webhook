/*
@File: server_test.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-28 20:39  Lucien     1.0         Init
*/
package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// requestJson 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func requestJson(t *testing.T, method string, uri string, param map[string]interface{}, router *gin.Engine) []byte {
	jsonByte, err := json.Marshal(param) // 将参数转化为json比特流
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte)) // 构造post请求，json数据以请求body的形式传递
	w := httptest.NewRecorder()                                        // 初始化响应
	router.ServeHTTP(w, req)                                           // 调用相应的handler接口
	result := w.Result()                                               // 提取响应
	body, err := ioutil.ReadAll(result.Body)                           // 读取响应body
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func set(t *testing.T, router *gin.Engine, Key string, Lang string, Content string, Password string) []byte {
	params := make(map[string]interface{})
	params[""] = Lang
	params["content"] = Content
	params["password"] = Password
	return requestJson(t, "POST", uri, params, router)
}
