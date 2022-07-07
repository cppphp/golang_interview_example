package main

//var定义块内，后面的定义可以使用前面定义的变量
var (
	size    = 1024
	maxSize = size * 2
)

func main() {
	println(size, maxSize)
}

/**
1024 2048
*/
