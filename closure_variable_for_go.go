package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			//goroutine的调度运行和i的赋值是并行的，i的值是随机的，相对来说i的赋值更快，结果中5的出现次数居多。
			fmt.Printf("%d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
$ GO111MODULE=off go run test.go
3
5
5
5
5

$ GO111MODULE=off go run test.go
5
3
5
5
5

$ GO111MODULE=off go run test.go
5
5
5
5
5

$ GO111MODULE=off go run test.go
2
5
5
5
5

$ GO111MODULE=off go run test.go
5
5
3
5
1

$ GO111MODULE=off go run test.go
5
5
5
3
5

$ GO111MODULE=off go run test.go
5
5
5
5
3
*/
