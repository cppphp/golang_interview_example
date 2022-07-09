package main

func main() {
	var a, b = 1, 2
	//3
	println(a | b)

	var c, d = 1.0, 2.0
	//invalid operation: c | d (operator | not defined on float64)
	println(c | d)
}

/**
invalid operation: c | d (operator | not defined on float64)
*/
