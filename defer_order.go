package main

import "fmt"

func main() {
	defer_call()
}

//堆栈，后进先出
func defer_call() {
	defer func() {
		fmt.Println(" 打印后 " )
	}()
	defer func() {
		fmt.Println(" 打印中 " )
	}()
	defer func() {
		fmt.Println(" 打印前 " )
	}()
	panic("触发异常")
}
