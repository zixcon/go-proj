package client

import (
	"fmt"
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
		fmt.Println(err)
		//panic(err)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	return req
}

func DoReqeust(req *http.Request) []byte {
	log.Println("GID:", util.GoID(),"http请求开始")
	start := time.Now()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	elapsed := time.Since(start)
	log.Println("GID:", util.GoID(),"http请求结束,总共耗时: ", elapsed)
	return body
}
