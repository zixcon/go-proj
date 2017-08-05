package main

import (
	"util/netease"
	"util"
	"runtime"
	"log"
)

func main() {

	// 这个函数设置的是Go语言跑几个线程。
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	// 这个函数返回当前有的CPU数。
	log.Println("GID:", util.GoID(), "CPUs:", runtime.NumCPU(), "Goroutines:", runtime.NumGoroutine())

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:9999", nil))
	//	pprof.Handler("")
	//}()

	url := "http://quotes.money.163.com/hs/service/diyrank.php"

	ch1 := make(chan string)
	ch2 := make(chan string)

	go netease.DealA(ch1, url)
	for {
		select {
		case url1 := <-ch1:
			go netease.DoAquoteCh(ch2, url1)
		case url2 := <-ch2:
			go netease.DealOne(url2)
		}
	}
}
