package main

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

//如果有命名的返回参数，return先给参数赋值，然后再执行defer语句可影响返回值，然后再返回
//如果无命名的返回参数，return决定了返回值，然后再执行defer语句时更改变量值不影响返回值，然后再返回

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

/**
4
1
3
*/
