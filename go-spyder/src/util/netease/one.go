package netease

import (
	"time"
	"log"
	"strconv"
	"util/client"
)

func WY_Header() map[string]string {
	header := map[string]string{
		"Host":                      "quotes.money.163.com",
		"Accept-Language":           "zh-CN,zh;q=0.8",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		//"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"User-Agent": RadomUA(),
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
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

func DealOne(url string) {
	season := [4]int{4, 3, 2, 1}
	year := time.Now().Year()
	to_year := time.Now().Year()
	loop := true
	for loop {
		for i := 0; i < 4; i++ {
			_, content := CallOne(year, season[i], url)
			if len(content) <= 0 && year < to_year {
				loop = false
				break
			}
		}
		year--
	}
}

func CallOne(year int, season int, url string) ([]string, map[string][]string) {
	log.Println()
	param := WY_Get_Param(year, season)
	req := client.GetRequest(url, param, WY_Header())
	bytes := client.DoReqeust(req)
	body := string(bytes)
	//log.Println(body)
	log.Println("请求结果处理开始")
	start := time.Now()
	title := HtmlTitle(body)
	content := HtmlContent(body)
	log.Println(title)
	log.Println(content)
	elapsed := time.Since(start)
	log.Println("请求结果处理结束,总共耗时: ", elapsed)
	return title, content
}

//func main() {
//	param := map[string]string{
//		"year":   "1900",
//		"season": "1",
//	}
//	url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
//	req := GetRequest(url, param)
//	body := DoReqeust(req)
//
//	title := HtmlTitle(body)
//	content := HtmlContent(body)
//	fmt.Println(title)
//	fmt.Println(content)
//}
