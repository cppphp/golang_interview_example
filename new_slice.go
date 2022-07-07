package main

import "fmt"

func main() {
	//new出来的是一个切片类型的指针；正确的方式是make出来一个切片才可以操作
	//list := make([]int, 0)
	list := new([]int)
	//first argument to append must be slice; have *[]int
	list = append(list, 1)
	fmt.Println(list)
}
