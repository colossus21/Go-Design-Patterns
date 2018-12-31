package main

import "fmt"

//This pattern can also be called Factory of Factory Pattern ie. Abstract Factory
//We have been making coffee for a long time, it's time to introduce Tea to our list
//Note that, Whether tea or we choose water, all of them can be categorized into Drinks category (call it a super class? this isn't Java though)

//So we have DRINKS -> COFFEE, TEA
//We will make 2 Types of Coffee this time -> Espresso, Mocha
//And, 2 Types of Tea -> Milk Tea, Green Tea

//We need a Factory to create Drinks -> DrinksMaker
//DrinksMaker have two subtypes -> TeaMaker, CoffeeMaker

//Coffee
type Coffee interface {
	GetRoastType() string
}

type Espresso struct {
	sugar int
	milk int
	roast string
}

type Mocha struct {
	sugar int
	milk int
	roast string
	chocolate int
}

//Espresso and Mocha must implement both interfaces - Drinks, Coffee
//Drinks -> GetSugarAmount, GetMilkAmount, Coffee -> GetRoastType

func (c *Espresso) GetSugarAmount() int {
	c.sugar = 20
	return c.sugar
}
func (c *Espresso) GetMilkAmount() int {
	c.sugar = 0
	return c.milk
}
func (c *Espresso) GetRoastType() string {
	c.roast = "Dark Roast"
	return c.roast
}

func (c *Mocha) GetSugarAmount() int {
	c.sugar = 50
	return c.sugar
}
func (c *Mocha) GetMilkAmount() int {
	c.milk = 30
	return c.milk
}
func (c *Mocha) GetRoastType() string {
	c.roast = "Medium-Dark Roast"
	return c.roast
}

//Tea
type Tea interface {
	GetTeaLeavesName() string
}

type MilkTea struct {
	sugar int
	milk int
	leavesName string
}

type GreenTea struct {
	sugar int
	milk int
	leavesName string
}

//MilkTea and GreenTea must implement both interfaces - Drinks, Tea
//Drinks -> GetSugarAmount, GetMilkAmount, Tea -> GetTeaLeavesName

func (t *MilkTea) GetSugarAmount() int {
	t.sugar = 60
	return t.sugar
}
func (t *MilkTea) GetMilkAmount() int {
	t.milk = 75
	return  t.milk
}
func (t *MilkTea) GetTeaLeavesName() string {
	t.leavesName = "Olong Tealeaves"
	return t.leavesName
}

func (t *GreenTea) GetSugarAmount() int {
	t.sugar = 10
	return t.sugar
}
func (t *GreenTea) GetMilkAmount() int {
	t.milk = 0
	return  t.milk
}
func (t *GreenTea) GetTeaLeavesName() string {
	t.leavesName = "Chinese Green Tealeaves"
	return t.leavesName
}


//Make Top-level Interfaces (Abstracts)

type Drink interface {
	GetSugarAmount() int
	GetMilkAmount() int
}

type DrinksMaker interface {
	MakeDrink(n int) Drink
}

//Define CONSTS

const (
	espresso = 1
	mocha = 2
	//We will have separate logic for these two, no worries!!
	greenTea = 1
	milkTea = 2
)


//Make it concrete: CoffeeMaker, TeaMaker

type CoffeeMaker struct {}

func (c CoffeeMaker) MakeDrink(n int) Drink {
	switch n {
	case espresso:
		coffee := new(Espresso)
		fmt.Println("Espresso Contents: Milk-",coffee.GetMilkAmount()," Sugar-",coffee.GetSugarAmount(), " Roast Type-",coffee.GetRoastType())
		return new(Espresso)
	case mocha:
		coffee := new(Mocha)
		fmt.Println("Mocha Contents: Milk-",coffee.GetMilkAmount()," Sugar-",coffee.GetSugarAmount(), " Roast Type-",coffee.GetRoastType())
		return new(Mocha)
	default:
		panic("Invalid Coffee Type!!")
	}
}

type TeaMaker struct {}

func (t TeaMaker) MakeDrink(n int) Drink {
	switch n {
	case milkTea:
		tea := new(MilkTea)
		fmt.Println("Milk Tea Contents: Milk-",tea.GetMilkAmount()," Sugar-",tea.GetSugarAmount(), " Tealeaves Name-",tea.GetTeaLeavesName())
		return tea
	case greenTea:
		tea := new(GreenTea)
		fmt.Println("Green Tea Contents: Milk-",tea.GetMilkAmount()," Sugar-",tea.GetSugarAmount(), " Tealeaves Name-",tea.GetTeaLeavesName())
		return tea
	default:
		panic("Invalid Tea Type!!")
	}
}

func main()  {
	tea1 := TeaMaker{}.MakeDrink(milkTea)
	tea2 := TeaMaker{}.MakeDrink(greenTea)
	coffee1 := CoffeeMaker{}.MakeDrink(espresso)
	coffee2 := CoffeeMaker{}.MakeDrink(mocha)
	fmt.Println("Total Milk Contents: ",tea1.GetMilkAmount()+tea2.GetMilkAmount()+coffee1.GetMilkAmount()+coffee2.GetMilkAmount())
	fmt.Println("Total Sugar Contents: ",tea1.GetSugarAmount()+tea2.GetSugarAmount()+coffee1.GetSugarAmount()+coffee2.GetSugarAmount())
}