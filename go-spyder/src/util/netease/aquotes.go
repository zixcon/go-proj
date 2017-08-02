package netease

import (
	"strconv"
	"log"
	"util/client"
	"encoding/json"
	"time"
	"github.com/golang/glog"
)

// A股 http://quotes.money.163.com/hs/service/diyrank.php
type QuotePage struct {
	Total     int `json:"total"`
	Pagecount int `json:"pagecount"`
	Quotes    []Quote `json:"list"`
}

type Quote struct {
	CODE   string
	NAME   string
	SNAME  string
	SYMBOL string
}

func json2struct(body []byte) (*QuotePage, error) {
	var quotePage *QuotePage
	err := json.Unmarshal(body, &quotePage)
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
	pageNo := 1
	pageSize := 24
	for pagecount := pageNo; pageNo <= pagecount; pageNo ++ {
		if pageNo == 1 {
			pagecount = doA(pageNo, pageSize, url)
		} else {
			go doA(pageNo, pageSize, url)
		}
	}
	time.Sleep(time.Second * 30)
}

func doA(pageNo int, pageSize int, url string) int {
	body := CallPage(pageNo, pageSize, url)
	//log.Println(string(body))
	log.Println(pageNo, "请求结果处理开始")
	glog.Infoln(pageNo, "请求结果处理开始")
	start := time.Now()
	quotePage, err := json2struct(body)
	if err != nil {
		log.Println(pageNo, "json 2 obj error")
	}
	resp,_ := json.Marshal(quotePage)
	log.Println(pageNo, string(resp))
	//for i := 0; i < pageSize; i++ {
	//	log.Println(quotePage.Quotes[i].NAME)
	//}

	urlArr := DoAquote(quotePage.Quotes)
	for i := 0; i < len(urlArr); i++ {
		DealOne(urlArr[i])
	}

	elapsed := time.Since(start)
	log.Println(pageNo, "请求结果处理结束,总共耗时: ", elapsed)
	return quotePage.Pagecount
}

func CallPage(pageNo int, pageSize int, url string) []byte {
	log.Println()
	param := WY_A_Get_Param(pageNo, pageSize)
	req := client.GetRequest(url, param, WY_A_Header())
	body := client.DoReqeust(req)
	//log.Println(body)
	return body
}
