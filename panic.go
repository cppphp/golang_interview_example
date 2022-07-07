package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		//覆盖之前的panic
		panic("defer panic")
	}()
	panic("panic")
}

/**
recover defer panic
*/
