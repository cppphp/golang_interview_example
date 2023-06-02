package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		fmt.Printf("for %d\n", i)
		wg.Add(1)
		go func() {
			//闭包使用的外部变量，且外部变量是传值类型，但其实是同一个变量，所以for打印的次数有时不满5次。
			//goroutine的调度运行和i的赋值是并行的，i的值是随机的，相对来说i的赋值更快，所以结果中go的i值大于10的居多。
			i = i + 5
			fmt.Printf("go  %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
$ GO111MODULE=off go run test.go
for 0
for 1
for 2
go  7
go  13
go  18

$ GO111MODULE=off go run test.go
for 0
for 1
for 2
for 3
for 4
go  10
go  15
go  20
go  25
go  30

$ GO111MODULE=off go run test.go
for 0
for 1
for 2
go  13
go  7
go  18

$ GO111MODULE=off go run test.go
for 0
for 1
for 2
for 3
go  14
go  19
go  8
go  24

$ GO111MODULE=off go run test.go
for 0
for 1
for 2
for 3
for 4
go  10
go  15
go  20
go  25
go  30

$ GO111MODULE=off go run test.go
for 0
for 1
for 2
for 3
for 4
go  10
go  15
go  20
go  30
go  25
*/
