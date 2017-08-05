package netease

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
	"util/client"
	"sync"
	"util"
)

// A股 http://quotes.money.163.com/hs/service/diyrank.php
type QuotePage struct {
	Total     int     `json:"total"`
	Pagecount int     `json:"pagecount"`
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

func DealA(ch chan<- string, url string) {
	pageNo := 1
	pageSize := 24
	var wg sync.WaitGroup
	for pagecount := pageNo; pageNo <= pagecount; pageNo++ {
		if pageNo == 1 {
			wg.Add(1)
			defer wg.Done()
			pagecount = doA(ch, pageNo, pageSize, url)
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()
				doA(ch, pageNo, pageSize, url)
			}()
		}
	}
	wg.Wait()
	close(ch)
	// time.Sleep(time.Second * 30)
}

func doA(ch chan<- string, pageNo int, pageSize int, url string) int {
	body := CallPage(pageNo, pageSize, url)
	if len(body) > 0 {
		log.Println("GID:", util.GoID(), "请求结果处理开始:", string(body))
		start := time.Now()
		quotePage, err := json2struct(body)
		if err != nil {
			log.Println("GID:", util.GoID(), "json 2 obj error")
			return -1
		}
		//resp, err := json.Marshal(quotePage)
		//if err != nil {
		//	log.Println("GID:", util.GoID(), "obj 2 json error")
		//}
		//log.Println("GID:", util.GoID(), "请求返回结果", string(resp))
		//for i := 0; i < pageSize; i++ {F
		//	log.Println(quotePage.Quotes[i].NAME)
		//}
		if quotePage != nil {
			DealAquoteCh(ch, quotePage.Quotes)
		}

		//urlArr := DoAquote(quotePage.Quotes)
		//for i := 0; i < len(urlArr); i++ {
		//	ch <- urlArr[i]
		//	// DealOne(urlArr[i])
		//}

		elapsed := time.Since(start)
		log.Println("GID:", util.GoID(), "请求结果处理结束,总共耗时: ", elapsed)
		return quotePage.Pagecount
	}
	return -1
}

func CallPage(pageNo int, pageSize int, url string) []byte {
	param := WY_A_Get_Param(pageNo, pageSize)
	req := client.GetRequest(url, param, WY_A_Header())
	body := client.DoReqeust(req)
	return body
}

// ------------------------

func DealAquoteCh(ch chan<- string, quotes []Quote) {
	for i := 0; i < len(quotes); i++ {
		quote := quotes[i]
		ch <- url + "/" + quote.CODE + ".html"
	}
}
