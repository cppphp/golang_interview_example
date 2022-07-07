package main

const cl = 100

var bl = 123

func main() {
	println(&bl, bl)
	//常量不能取地址
	println(&cl, cl)
}

/**
invalid operation: cannot take address of cl (untyped int constant 100)
*/
