package main

import (
	"fmt"
)

type student struct {
	Name string
	Age int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//stu是传值变量，不是原地址，且range复用此变量
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	stus[2].Name = "newname"

	for i := range m {
		fmt.Printf("%v - %v\n", i, m[i])
	}

	for i := range stus {
		fmt.Printf("%v - %v\n", i, stus[i])
	}
}

func main() {
	paseStudent()
}
