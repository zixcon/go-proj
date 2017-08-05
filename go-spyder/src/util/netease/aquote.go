package netease

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"util/client"
	"util"
)

var url = "http://quotes.money.163.com"

func WY_Aquote_Header() map[string]string {
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

func DealQuote(quotes []Quote) []string {
	arr := make([]string, len(quotes))
	for i := 0; i < len(quotes); i++ {
		quote := quotes[i]
		arr[i] = url + "/" + quote.CODE + ".html"
	}
	return arr
}

func forOneUrl(body string) string {
	//log.Println(body)
	bodyReader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Println("GID:", util.GoID(), err)
	}
	var href string
	doc.Find("#menuCont").Find(".submenu_cont").Find(".sub_menu").EachWithBreak(func(i int, s *goquery.Selection) bool {
		s.Find("ul").Find("li").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if strings.EqualFold("历史交易数据", s.Text()) {
				href, _ = s.Find("a").Attr("href")
				return false
			}
			return true
		})
		return true
	})
	log.Println("GID:", util.GoID(), href)
	return url + href
}

func CallAquote(url string) []byte {
	req := client.GetRequest(url, nil, WY_Aquote_Header())
	body := client.DoReqeust(req)
	//log.Println(body)
	return body
}

// 历史交易数据url
func DoAquote(quotes []Quote) []string {
	var arr = make([]string, len(quotes))
	aquoteArr := DealQuote(quotes)
	for i := 0; i < len(aquoteArr); i++ {
		body := CallAquote(aquoteArr[i])
		url := forOneUrl(string(body))
		log.Println("GID:", util.GoID(), url)
		arr[i] = url
	}
	return arr
}

// ------------------------

func DoAquoteCh(ch chan<- string, srcUrl string) {
	body := CallAquote(srcUrl)
	if len(body) > 0 {
		url := forOneUrl(string(body))
		ch <- url
	}
}
