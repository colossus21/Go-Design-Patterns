package main

import (
	"fmt"
	"strconv"
)

//In this pattern, you have a Factory that returns a Function based on the provided information
//Multiple objects of same type can have a method that is to be returned
//But, the methods are implemented differently
//In this case, we have multiple structs having the same interface

type Cappuccino struct {
	name string
	ingredients string
	foamAmount int
}

type Espresso struct {
	name string
	ingredients string
	sugarAmount int
}

type Mocha struct {
	name string
	ingredients string
	chocolateAmount int
}

// Let's create an interface for the common methods
type CoffeeMakingProcess interface {
	MakeCoffee(extraStaff int) string
}
//--End--

func (e *Cappuccino) MakeCoffee(extraStaff int) string {
	fmt.Println("Cappuccino makes me stronger!!")
	e.name = "Cappuccino"
	e.ingredients = "Espresso, Lots of Milk Foam!"
	e.foamAmount = extraStaff
	return "Order served: A Cup of " + e.name + " having " + e.ingredients + " and " + strconv.Itoa(e.foamAmount) + " cubic cm foam!"
}

func (e *Espresso) MakeCoffee(extraStaff int) string {
	fmt.Println("An Espresso a day keeps the doctor away!!")
	e.name = "Espresso"
	e.ingredients = "One shot of Espresso!"
	e.sugarAmount = extraStaff
	return "Order served: A Cup of " + e.name + " having " + e.ingredients + " and " + strconv.Itoa(e.sugarAmount) + " gm of sugar"
}

func (e *Mocha) MakeCoffee(extraStaff int) string {
	fmt.Println("Here you go, a cup of refreshing Mocha!!")
	e.name="Mocha"
	e.ingredients="Espresso, Lots of Chocolate!"
	e.chocolateAmount = 45
	return "Order served: A Cup of "+ e.name+ " having "+ e.ingredients+ " and "+ strconv.Itoa(e.chocolateAmount)+ " pieces chocolate chips!"
}

const (
	cappuccino = 1
	mocha = 2
	espresso = 3
)

//The Coffee Types are implementing CoffeeMakingProcess because, we have implemented the Interface in all of them

func ServeCoffee(n int) CoffeeMakingProcess {
	switch n {
	case cappuccino:
		return new(Cappuccino)
	case mocha:
		return new(Mocha)
	case espresso:
		return new(Espresso)
	default:
		panic("Invalid Coffee Type!!")
	}
}

func main() {
	order1 := ServeCoffee(cappuccino)
	fmt.Println(order1.MakeCoffee(69))
	fmt.Println()

	order2 := ServeCoffee(mocha)
	fmt.Println(order2.MakeCoffee(420))
	fmt.Println()

	order3 := ServeCoffee(espresso)
	fmt.Println(order3.MakeCoffee(420+69))
	fmt.Println()

	//n=4 is invalid
	order4 := ServeCoffee(4)
	order4.MakeCoffee(911)
}

