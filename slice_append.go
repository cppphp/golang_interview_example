package main

import "fmt"

func main() {
	a := make([]int, 0, 3)
	//append返回的是个新的slice结构体，虽然array指针是相同的，但len值却不同。
	b := append(a, 1)
	//a的len依然是0，所以打印[]
	fmt.Printf("%v\n", a)
	//b的len是1，含有一个元素。
	fmt.Printf("%v\n", b)
}

/**
[]
[1]
*/
