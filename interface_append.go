package main

import "fmt"

func main() {
	s1 := []interface{}{}
	s2 := []int{1, 2}
	//s2整体看作interface{}
	s1 = append(s1, s2)
	//1、2看作interface{}
	s1 = append(s1, 1, 2)
	////s2...类型不匹配
	////cannot use s2 (type []int) as type []interface {} in append
	//s1 = append(s1, s2...)
	s3 := []interface{}{1, 2}
	s1 = append(s1, s3...)
	fmt.Println(s1)
}

/**
[[1 2] 1 2 1 2]
*/
