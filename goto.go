package main

func main() {
	for i := 0;i < 10;i++ {
loop:
		println(i)
	}
	//不可以从循环外部goto到循环内部
	goto loop
}

/**
goto loop jumps into block starting at 4:26
*/
