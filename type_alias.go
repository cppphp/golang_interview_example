package main

func main() {
	type MyInt1 = int //只是别名
	type MyInt2 int   //定义新类型
	var i int = 1
	//ok
	var i1 MyInt1 = i
	println(i1)
	//不同类型不能赋值
	var i2 MyInt2 = i
	println(i2)
}

/**
cannot use i (type int) as type MyInt2 in assignment
*/
