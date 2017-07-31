package netease

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func HtmlTitle(body string) []string {
	var titles []string
	bodyReader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(".table_bg001 ").EachWithBreak(func(i int, s *goquery.Selection) bool {
		// 标题
		th := s.Find("thead").Find("th")
		log.Println("标题Size:" , th.Size())
		arr := make([]string, th.Size())
		th.EachWithBreak(func(j int, g *goquery.Selection) bool {
			th_title := g.Text()
			arr[j] = th_title
			//if (j >= th.Size()) {
			//	return false
			//} else {
			//	return true
			//}
			return true
		})
		titles = arr
		return false
	})
	return titles
}

func HtmlContent(body string) map[string][]string {
	var content map[string][]string
	var key string
	bodyReader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Fatalln(err)
	}
	/* 创建集合 */
	content = make(map[string][]string)
	doc.Find(".table_bg001 ").EachWithBreak(func(i int, s *goquery.Selection) bool{
		// 内容
		s.Find("tr").EachWithBreak(func(j int, g *goquery.Selection) bool{
			td := g.Find("td")
			size := td.Size()
			//log.Println("内容size:" , size)
			body := make([]string, size )
			td.Each(func(k int, h *goquery.Selection) {
				td_body := h.Text()
				if k == 0 {
					key = td_body
				} else {
					body[k-1] = td_body
				}
				//fmt.Print(td_body + " ")
			})
			if len(key) > 0 && len(body) >0 {
				content[key] = body
			}
			//fmt.Println()
			return true
		})
		return false
	})
	return content
}
