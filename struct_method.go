package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u user) setname(n string) {
	u.name = n
}

func (u *user) setage(a int) {
	u.age = a
}

func main() {
	var u1 = user{
		name: "a",
		age:  18,
	}
	fmt.Println(u1)
	//{a 18}
	//struct是值类型，以值类型为对象的方法，方法内不能修改对象本身
	u1.setname("b")
	//以指针类型为对象的方法，方法内可以修改对象本身
	u1.setage(20)
	fmt.Println(u1)
	//{a 20}
}
