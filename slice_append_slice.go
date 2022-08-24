package main

import "fmt"

func main() {
	a := [][]int{}
	b := []int{1,2}
	//slice是引用类型，所以此处的b是个引用
	c := append(a, b)
	//将b展开为int后再append，b就不再是引用了
	d := append(a, append([]int{}, b...))
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	b[0] = 3
	//c含有对b的引用，所以b的值会跟着变化
	fmt.Printf("%v\n", c)
	//d不含对b的引用，值不变化
	fmt.Printf("%v\n", d)
}

/**
[[1 2]]
[[1 2]]
[[3 2]]
[[1 2]]
*/
