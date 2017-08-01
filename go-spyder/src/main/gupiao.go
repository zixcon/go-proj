package main

import (
	"util/netease"
)

func main() {

	//ua := netease.RadomUA()
	//log.Println(ua)

	//url := "http://quotes.money.163.com/trade/lsjysj_600570.html"
	//netease.DealOne(url)

	url := "http://quotes.money.163.com/hs/service/diyrank.php"
	netease.DealA(url)
}
