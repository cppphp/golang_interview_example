package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("nil interface")
		return
	}
	fmt.Println("non-nil interface")
}

func main() {
	var x *int = nil
	Foo(x)
	//类型和值都为nil才是nil
	var y interface{} = nil
	Foo(y)
	var z interface{} = 1
	Foo(z)
}

/**
non-nil interface
nil interface
non-nil interface
*/
