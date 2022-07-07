package main

////ok
//func GetValue() interface{} {
////error
func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	//cannot type switch on non-interface value i (type int)
	//类型断言需要是接口类型才可以
	switch i.(type) {
		case int:
			println("int")
		case string:
			println("string")
		case interface{}:
			println("interface")
		default:
			println("unknown")
	}
}
