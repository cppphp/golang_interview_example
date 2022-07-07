package main

import "fmt"

type User struct {
	id int
	name string
}

type Manager struct {
	User
	title string
}

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	//多态的实现
	m := Manager{User{1, "Tom"}, "Administrator"}
	//当Manager.ToString未定义时，此处打印：User: 0xc00005c300, &{1 Tom}
	//当Manager.ToString已定义时，此处打印：Manager: 0xc00005c300, &{{1 Tom} Administrator}
	fmt.Println(m.ToString())
	//不管Manager.ToString是否定义，此处都打印：User: 0xc00005c300, &{1 Tom}
	fmt.Println(m.User.ToString())
}
