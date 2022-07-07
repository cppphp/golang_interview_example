package main

import"fmt"

type T1 struct { }

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1

//匿名成员只能同类型有一个，不能有两个T1，虽然通过type绕过了限制，但下面出错了
type MyStruct struct {
	T1
	T2
}

func main() {
	my := MyStruct{}
	//error: 模棱两可的选择器 my.m1
	my.m1()
}

/**
ambiguous selector my.m1
*/
