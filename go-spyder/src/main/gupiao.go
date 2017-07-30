package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
	"strings"
	"bytes"
)

// 爬取 xxxx 历年每天的交易数据
// http://quotes.money.163.com/trade/lsjysj_600570.html?year=2016&season=3

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

func stringJoin(arr []string, seprator string) string {
	str := bytes.Buffer{}
	for _, value := range arr {
		str.WriteString(value)
		str.WriteString(seprator)
	}
	return str.String()
}

func GetRequest(url string, param map[string]string) *http.Request {
	var reqUrl = url
	var paramStr = make([]string, 0, len(param))
	var index int16
	for key, value := range param {
		paramStr = append(paramStr, key+"="+value)
		//getParam[index] = key + "=" + value
		index ++
	}
	urlStr := reqUrl + "?" + strings.Join(paramStr, "&")
	fmt.Println("请求地址：", urlStr)
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
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func parse_title(htmlBody string) {
	body := unpackHtml(htmlBody)
	fmt.Println("返回内容：", body)
	pattern_title := `<tr class="dbrow">(.*?)</tr>`
	rp_title := regexp.MustCompile(pattern_title)
	find_title := rp_title.FindAllStringSubmatch(body, -1)
	fmt.Println(find_title)
}

func unpackHtml(html string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = re.ReplaceAllString(html, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = re.ReplaceAllString(html, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	//re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	//html = re.ReplaceAllString(html, "\n")

	//去除连续的换行符
	//re, _ = regexp.Compile("\\s{2,}")
	//html = re.ReplaceAllString(html, "\n")
	return html
}

func main() {
	param := map[string]string{
		"year":   "2016",
		"season": "1",
	}
	url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
	req := GetRequest(url, param)
	body := DoReqeust(req)
	parse_title(body)

}
