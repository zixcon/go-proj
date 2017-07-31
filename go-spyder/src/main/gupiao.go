package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"util/netease"
	"time"
	"strconv"
)

func WY_Header() map[string]string {
	header := map[string]string{
		"Host":                      "quotes.money.163.com",
		"Accept-Language":           "zh-CN,zh;q=0.8",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	return header
}

func WY_Get_Param(year int, season int) map[string]string {
	param := map[string]string{
		"year":   strconv.Itoa(year),
		"season": strconv.Itoa(season),
	}
	return param
}

func GetRequest(url string, param map[string]string) *http.Request {
	var reqUrl = url
	var paramStr = make([]string, 0, len(param))
	var index int16
	for key, value := range param {
		paramStr = append(paramStr, key+"="+value)
		//getParam[index] = key + "=" + value
		index++
	}
	urlStr := reqUrl + "?" + strings.Join(paramStr, "&")
	log.Println("请求地址：", urlStr)
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	for key, value := range WY_Header() {
		req.Header.Add(key, value)
	}
	return req
}

func DoReqeust(req *http.Request) string {
	log.Println("http请求开始")
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
	log.Println("http请求结束,总共耗时: ", elapsed)
	return string(body)
}

func dealOne(url string) {
	season := [4]int{4, 3, 2, 1}
	year := time.Now().Year()
	to_year := time.Now().Year()
	loop := true
	for loop {
		log.Println()
		for i := 0; i < 4; i++ {
			param := WY_Get_Param(year, season[i])
			req := GetRequest(url, param)
			body := DoReqeust(req)

			log.Println("请求结果处理开始")
			start := time.Now()
			title := netease.HtmlTitle(body)
			content := netease.HtmlContent(body)
			log.Println(title)
			log.Println(content)
			elapsed := time.Since(start)
			log.Println("请求结果处理结束,总共耗时: ", elapsed)
			if len(content) <= 0 && year < to_year {
				loop = false
				break
			}
		}
		year--
	}
}

func main() {
	//param := map[string]string{
	//	"year":   "1900",
	//	"season": "1",
	//}
	//url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
	//req := GetRequest(url, param)
	//body := DoReqeust(req)
	//
	//title := wangyi.HtmlTitle(body)
	//content := wangyi.HtmlContent(body)
	//fmt.Println(title)
	//fmt.Println(content)

	url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
	dealOne(url)
}
