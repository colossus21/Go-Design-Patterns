package main

import "fmt"

//Everything is decoupled - abstractions and concretes / signatures and implementations

//Level 1 Abstraction, IDrinks -> Tea, Coffee

type IDrinks interface {
	Make()
}

type Tea struct {
	sugar int
	milk int
	tealeaves int
}

func (t *Tea) Make() {
	fmt.Println("Boil water !!")
	fmt.Println("Add Sugar, amount: ",t.sugar)
	fmt.Println("Add Milk, amount: ",t.milk)
	fmt.Println("Add Tealeaves, amount: ",t.tealeaves)
	fmt.Println("Pour the mix in a cup after 10 minutes!")
}

type Coffee struct {
	sugar int
	liquidMilk int
	coffee int
}

func (c *Coffee) Make() {
	fmt.Println("Boil milk, amount: ", c.liquidMilk)
	fmt.Println("Add Coffee, amount: ",c.coffee)
	fmt.Println("Add Sugar, amount: ",c.sugar)
	fmt.Println("Pour the mix in a cup after 1 hour!")
}

//Level 2 Implementer Abstractions

type IDrinksMaker interface {
	SetDrink(drink IDrinks) IDrinksMaker
	RemoveDrink() IDrinksMaker
	makeDrink() IDrinksMaker
}

//Takes any DRINK and make it.

type DrinksMaker struct {
	d IDrinks
}

func (d *DrinksMaker) SetDrink(drink IDrinks) IDrinksMaker {
	d.d = drink
	return d
}

func (d *DrinksMaker) RemoveDrink() IDrinksMaker {
	d.d = nil
	return d
}

func (d *DrinksMaker) makeDrink() IDrinksMaker {
	d.d.Make()
	return d
}

func main() {
	tea := new(Tea)
	coffee := new(Coffee)

	tea.sugar = 10
	tea.milk = 20
	tea.tealeaves = 30

	coffee.sugar = 40
	coffee.liquidMilk = 50
	coffee.coffee = 60

	machine := new(DrinksMaker)

	//The magic of bridge pattern
	machine.SetDrink(tea).makeDrink().RemoveDrink().SetDrink(coffee).makeDrink()
}