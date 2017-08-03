package util

import (
	"runtime"
	"strings"
	"strconv"
	"log"
	"sync"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		log.Println("cannot get goroutine id: %v", err)
	}
	return id
}

func test() {
	log.Println("main", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println(i, GoID())
		}()
	}
	wg.Wait()
}
