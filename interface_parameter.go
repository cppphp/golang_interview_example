package main

import (
	"fmt"
)

func main(){
	var v int = 1
	var vv = []int{1,2}

	t1(v, vv)

	////cannot use vv (type []int) as type []interface {} in argument to t2
	//t2(v, vv)

	var ii = []interface{}{1,2}
	t2(v, ii)
}

//interface{}参数类型可以接收任何参数
func t1(v interface{}, vv interface{}) {
	fmt.Printf("%v, %v\n", v, vv)

	////invalid operation: vv[0] (type interface {} does not support indexing)
	//fmt.Printf("%v\n", vv[0])
}

//[]interface{}参数类型只能接收[]interface{}类型参数
func t2(v interface{}, vv []interface{}) {
	fmt.Printf("%v, %v\n", v, vv)
	fmt.Printf("%v\n", vv[0])
}

/**
1, [1 2]
1, [1 2]
1
*/
