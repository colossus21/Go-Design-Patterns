package main

import "fmt"

type TemplateMethods interface {
	Operation1() int
	Operation2() int
	Operation3() string
}

func TemplateMethod(T TemplateMethods) {
	fmt.Println("Calculating for",T.Operation1(),"and",T.Operation2())
	fmt.Print("Result: ")
	if(T.Operation3() == "+") {
		fmt.Print(T.Operation1()+T.Operation2())
	} else if (T.Operation3()=="-") {
		fmt.Print(T.Operation1()-T.Operation2())
	} else if (T.Operation3()=="x") {
		fmt.Print(T.Operation1()*T.Operation2())
	} else if (T.Operation3()=="/") {
		fmt.Print(T.Operation1()/T.Operation2())
	} else {
		fmt.Print("Result is unknown")
	}
}

//Implementation

type Adder struct {}

func (a *Adder) Operation1() int {
	return 100
}

func (a *Adder) Operation2() int {
	return 450
}

func (a *Adder) Operation3() string {
	return "+"
}


func main() {
	A := Adder{}
	TemplateMethod(&A)
}