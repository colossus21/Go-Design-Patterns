package main

import (
	"fmt"
)

//Let's recreate the shit we have done in the creational pattern

//Simple Factory Pattern is self-explanatory, you might not want to call it a pattern if you don't want to
//You have a Factory that returns an Object based on the provided information

type Recipe struct {
	name string
	ingredients string
}

func NewCappuccino() Recipe {
	e := Recipe{"Cappuccino", "1 Shot of espresso, Steamed Milk, Milk Foam, Sprinkled Chocolate"}
	fmt.Println("Making a cup of", e.name, " which contains ", e.ingredients, "!")
	return e
}

func NewMocha() Recipe {
	e := Recipe{"Mocha", "1 Shot of espresso, 1 Spoon of Chocolate Powder, Steamed Milk, Sprinkled Chocolate"}
	fmt.Println("Making a cup of", e.name, " which contains ", e.ingredients, "!")
	return e
}

func NewEspresso() Recipe {
	e := Recipe{"Espresso", "1 Shot of espresso in an espresso cup"}
	fmt.Println("Making a cup of", e.name, " which contains ", e.ingredients, "!")
	return e
}

const (
	cappuccino = 1
	mocha = 2
	espresso = 3
)

func MakeCoffee(n int) Recipe {
	switch n {
	case cappuccino:
		return NewCappuccino()
	case mocha:
		return NewMocha()
	case espresso:
		return NewEspresso()
	default:
		panic("Invalid Coffee Type!!")
	}
}

func main() {
	order1 := MakeCoffee(cappuccino)
	order2 := MakeCoffee(mocha)
	order3 := MakeCoffee(espresso)
	order4 := MakeCoffee(4)
	//The following line will not work
	fmt.Println("Orders placed: ", order1.name, order2.name, order3.name, order4.name)
}

