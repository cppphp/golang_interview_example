package main

import "fmt"

func main() {
	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//结构体是值类型，可以进行相等比较，但不能进行不等比较
	//invalid operation: sn1 > sn2 (operator > not defined on struct)
	if sn1 > sn2 {
		fmt.Println("sn1 > sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//map是引用类型，不能进行相等比较，被结构体包含后，结构体也变得不能进行相等比较
	//invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
