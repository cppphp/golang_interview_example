package main

import "fmt"

//iota中间出现指定的字符串或数字，接下来就会重复上面的指定，如果恢复iota则会继续常规的计数，包括上面指定的也会包含在计数内
const (
	x = iota
	y
	z = "zz"
	a
	b
	c = iota
	d
	e = 1
	f
	g = iota
	h
)

func main() {
	fmt.Println(x, y, z, a, b, c, d, e, f, g, h)
}

/**
0 1 zz zz zz 5 6 1 1 9 10
*/
