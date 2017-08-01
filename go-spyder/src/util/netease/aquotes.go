package netease

import (
	"strconv"
	"log"
	"util/client"
	"encoding/json"
	"time"
)

// A股 http://quotes.money.163.com/hs/service/diyrank.php
type QuotePage struct {
	total     int
	pagecount int
	quotes    []Quote
}

type Quote struct {
	CODE   string
	NAME   string
	SNAME  string
	SYMBOL string
}

func json2struct(body []byte) (*QuotePage, error) {
	quotePage := &QuotePage{}
	err := json.Unmarshal(body, quotePage)
	return quotePage, err
}

func WY_A_Header() map[string]string {
	header := map[string]string{
		"Host":                      "quotes.money.163.com",
		"Accept-Language":           "zh-CN,zh;q=0.8",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                RadomUA(),
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	return header
}

func WY_A_Get_Param(pageNo int, pageSize int) map[string]string {
	param := map[string]string{
		"page":   strconv.Itoa(pageNo),
		"host":   "http://quotes.money.163.com/hs/service/diyrank.php",
		"query":  "STYPE:EQA",
		"fields": "NO,SYMBOL,NAME,PRICE,PERCENT,UPDOWN,FIVE_MINUTE,OPEN,YESTCLOSE,HIGH,LOW,VOLUME,TURNOVER,HS,LB,WB,ZF,PE,MCAP,TCAP,MFSUM,MFRATIO.MFRATIO2,MFRATIO.MFRATIO10,SNAME,CODE,ANNOUNMT,UVSNEWS",
		"sort":   "PERCENT",
		"order":  "asc",
		"type":   "query",
		"count":  strconv.Itoa(pageSize),
	}
	return param
}

func DealA(url string) {
	pagecount := 1
	pageNo := 1
	pageSize := 24
	for ; pageNo < pagecount; pageNo ++ {
		body := CallPage(pageNo, pageSize, url)
		log.Println(string(body))
		log.Println("请求结果处理开始")
		start := time.Now()
		quotePage, err := json2struct(body)
		if err != nil {
			log.Println("json 2 obj error")
		}
		pagecount = quotePage.pagecount
		elapsed := time.Since(start)
		log.Println("请求结果处理结束,总共耗时: ", elapsed)
	}
}

func CallPage(pageNo int, pageSize int, url string) []byte {
	log.Println()
	param := WY_A_Get_Param(pageNo, pageSize)
	req := client.GetRequest(url, param, WY_A_Header())
	body := client.DoReqeust(req)
	//log.Println(body)
	return body
}
