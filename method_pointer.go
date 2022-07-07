package main

import "fmt"

type Data struct{
	x int
}

//不支持重载，两个函数不可以重名

//传值调用，每次调用地址不同
func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}

//指针传递，每次调用指针值不变
func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest() // ValueTest(d)
	d.ValueTest() // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest() // ValueTest(*p)
	p.ValueTest() // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
	p.PointerTest() // PointerTest(p)
}

/*
Data:    0xc0000100a0
Value:   0xc0000100c0
Value:   0xc0000100c8
Pointer: 0xc0000100a0
Pointer: 0xc0000100a0
Value:   0xc0000100d0
Value:   0xc0000100d8
Pointer: 0xc0000100a0
Pointer: 0xc0000100a0
*/
