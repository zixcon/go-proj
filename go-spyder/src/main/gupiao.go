package main

import (
	"util/netease"
	"log"
)

func main() {

	ua := netease.RadomUA()
	log.Println(ua)
	//url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
	//netease.DealOne(url)
}
