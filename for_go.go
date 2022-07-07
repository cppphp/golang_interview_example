package main

import (
	"fmt"
	"runtime"
	"sync"
)

// A输出的都是A10；
// B输出的是无序的B0-B9
func main() {
	//保证协程和主协程串行调度
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		//闭包引用外部变量，i最终是10
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i:= 0; i < 10; i++ {
		//传值调用，i的值为调用时参数的值
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
