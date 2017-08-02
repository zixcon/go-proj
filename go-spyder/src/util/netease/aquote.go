package netease

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
	"log"
)

var url = "http://quotes.money.163.com/"

func DealQuote(quotes []*Quote) {
	for i := 0; i < quotes; i++ {
		quote := quotes[i]
		url + quote.CODE + ".html"
	}
}

func forOneUrl(body string) string {
	bodyReader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Fatalln(err)
	}
	var href string
	doc.Find(".area").Find(".title_01").EachWithBreak(func(i int, s *goquery.Selection) bool {
		s.Find("ul").Find("li").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if strings.EqualFold("历史交易数据", s.Text()) {
				href, _ = s.Find("a").Attr("href")
				return false
			}
			return true
		})
		return false
	})
	log.Println(href)
	return href
}
