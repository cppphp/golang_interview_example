package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

func main() {
	/*
	我们看到当把一个具体类型赋值给接口之后，这个接口一定不是nil，因为接口的动态类型保存着具体类型的type，
	但是接口的动态值可能是nil，所以只要有一个不为nil，那么接口就不是nil，这个大家要在以后的代码中谨慎一些。
	具体大家可以看下interface的源码实现就知道为什么了。
	 */
	t := live()
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", nil, nil)
	if t == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}

/**
B
*/
