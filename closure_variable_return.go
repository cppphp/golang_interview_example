package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			//引用的同一个变量
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

/**
0xc000010038 2
0xc000010038 2
*/
