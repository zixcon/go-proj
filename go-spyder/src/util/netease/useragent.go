package netease

import (
	"io/ioutil"
	"strings"
	"math/rand"
	"log"
	"runtime"
	"path/filepath"
	"sync"
)

var once sync.Once
var uaArr []string

// 获取调用者的当前文件DIR
//Get the caller now directory
func CurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func uaInit() []string {
	file := CurrentDir() + "/../.." + "/conf/ua.txt"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(file, "load error")
	}
	ua := string(data)
	//log.Println(ua)
	uaArr := strings.Split(ua, "\n")
	return uaArr
}

func onces() {
	uaArr = uaInit()
	log.Println("ua init")
}

func RadomUA() string {
	once.Do(onces)
	length := len(uaArr)
	if length <= 0 {
		return "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36"
	}
	return uaArr[rand.Intn(length)]
}
