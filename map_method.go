package main

import (
	"fmt"
)

type user map[string]interface{}

func (u user) setname(n string) {
	u["name"] = n
}

func (u *user) setage(a int) {
	(*u)["age"] = a
}

func main() {
	var u1 = user{"name":"a", "age":18}
	fmt.Println(u1)
	//map[age:18 name:a]
	//map是引用类型，以map类型为对象的方法，方法内可以修改对象本身
	u1.setname("b")
	//以指针类型为对象的方法，方法内可以修改对象本身
	u1.setage(20)
	fmt.Println(u1)
	//map[age:20 name:b]
}
