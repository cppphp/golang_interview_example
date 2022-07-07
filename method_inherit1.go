package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	//此处只能调用People的方法，如果People未定义此方法，则编译不通过。
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teachershowB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

/**
showA
showB
 */
