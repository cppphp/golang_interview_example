package main

import "fmt"

// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length, so make([]int, 0, 10) allocates a slice of length 0 and
//	capacity 10.
//	Map: An initial allocation is made according to the size but the
//	resulting map has length 0. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
// func make(Type, size IntegerType) Type

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
// func new(Type) *Type

func main() {
	//make创建语法，make([]int, len, cap)
	//s = []int{0,4:0} = [0,0,0,0,0]
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//[0 0 0 0 0 1 2 3]
}
