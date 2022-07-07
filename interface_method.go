package main

import "fmt"

type People interface {
	Speak(string) string
}

type User interface {
	Say(string) string
}

type Stduent struct{}

//只有 var peo People = &Stduent{} 可以调用
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "golang" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//用不用指针都可以调用
func (stu Stduent) Say(think string) (talk string) {
	if think == "golang" {
		talk = "You are a bad boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	think := "golang"
	////cannot use Stduent literal (type Stduent) as type People in assignment:
    ////    Stduent does not implement People (Speak method has pointer receiver)
	//var peo1 People = Stduent{}
	//fmt.Println(peo1.Speak(think))
	var peo2 People = &Stduent{}
	fmt.Println(peo2.Speak(think))
	var user1 User = Stduent{}
	fmt.Println(user1.Say(think))
	var user2 User = &Stduent{}
	fmt.Println(user2.Say(think))
}

/**
You are a good boy
You are a bad boy
You are a bad boy
*/
