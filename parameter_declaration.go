package main

////ok
//func Mui(x, y int) (sum int, err error) {
////syntax error: mixed named and unnamed function parameters
func Mui(x, y int) (sum int, error) {
	return x+y, nil
}
