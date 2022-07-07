package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	var (
		intChan = make(chan int, 1)
		stringChan = make(chan string, 1)
	)
	intChan <- 1
	stringChan <- "hello"
	//所有chan都已经有数据，case不需要依赖唤醒顺序
	//都已经ready的情况下，case执行顺序依赖于pollorder对case顺序的随机打乱
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
	//有时候1，有时候panic
}
