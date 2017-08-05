package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"util"
)

func GetRequest(url string, param map[string]string, header map[string]string) *http.Request {
	var urlStr = url
	if param != nil {
		var paramStr = make([]string, 0, len(param))
		var index int16
		for key, value := range param {
			paramStr = append(paramStr, key+"="+value)
			//getParam[index] = key + "=" + value
			index++
		}
		urlStr = url + "?" + strings.Join(paramStr, "&")
	}
	log.Println("GID:", util.GoID(), "请求地址：", urlStr)
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Println("GID:", util.GoID(), "请求地址组装失败", err)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	return req
}

func DoReqeust(req *http.Request) []byte {
	var body []byte
	log.Println("GID:", util.GoID(), "http请求开始")
	start := time.Now()
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("GID:", util.GoID(), "http请求失败", err.Error())
	} else {
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("GID:", util.GoID(), "http请求返回结果读取失败", err.Error())
		}
		elapsed := time.Since(start)
		log.Println("GID:", util.GoID(), "http请求结束,总共耗时: ", elapsed)
	}
	return body
}
