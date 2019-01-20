package main

import "fmt"

//EspressoMaker

type IEspressoMaker interface {
	createShots(n int) IEspressoMaker
}

type EspressoMaker struct {
	capacity float32
	litresPerShot float32
}

func (b *EspressoMaker) createShots(n int) IEspressoMaker {
	amount := float32(n)*b.litresPerShot
	if (amount > b.capacity) {
		panic("Espresso out of water!!")
	}
	fmt.Println("Creating", n, "shots of Espresso!")
	b.capacity -= amount
	return b
}

//CoffeeRoaster

type ICoffeeRoaster interface {
	roastCoffee(roastType string, amount float32) ICoffeeRoaster
}

type CoffeeRoaster struct {
	capacity float32
}

func (r *CoffeeRoaster) roastCoffee(roastType string, amount float32) ICoffeeRoaster {
	if(amount > r.capacity) {
		panic("Not enough coffee beans!!")
	}
	fmt.Println(roastType,"Roasting", amount, "mg of beans!")
	r.capacity -= amount
	return r
}

//Ingredients Adder

type IIngredients interface {
	addIngredients(args ...string) IIngredients
}

type Ingredients struct {}

func (i *Ingredients) addIngredients(args ...string) IIngredients {
	for _,v := range args {
		fmt.Println("Adding",v)
	}
	return i
}

//CoffeeMaker Facade

type ICoffeeMaker interface {
	MakeLatte() ICoffeeMaker
	MakeCappuccino() ICoffeeMaker
	MakeVienna() ICoffeeMaker
}

type CoffeeMaker struct {
	espresso EspressoMaker
	roaster CoffeeRoaster
	ingredients Ingredients
}

func (m *CoffeeMaker) MakeLatte() ICoffeeMaker {
	fmt.Println("Making Latte...")
	m.roaster.roastCoffee("Medium", 50)
	m.espresso.createShots(2)
	m.ingredients.addIngredients("Sugar", "Steamed Milk","Chocolate Sprinkles")
	fmt.Println("Latte Made!")
	fmt.Println()
	return m
}

func (m *CoffeeMaker) MakeCappuccino() ICoffeeMaker {
	fmt.Println("Making Cappuccino...")
	m.roaster.roastCoffee("Dark", 50)
	m.espresso.createShots(1)
	m.ingredients.addIngredients("Sugar", "Steamed Milk","Cream","Syrup")
	fmt.Println("Cappuccino Made!")
	fmt.Println()
	return m
}

func (m *CoffeeMaker) MakeVienna() ICoffeeMaker {
	fmt.Println("Making Vienna...")
	m.roaster.roastCoffee("Light", 50)
	m.espresso.createShots(3)
	m.ingredients.addIngredients("Whipped Cream","Condensed Milk")
	fmt.Println("Vienna Made!")
	fmt.Println()
	return m
}

func (m *CoffeeMaker) SetMachine(e *EspressoMaker, r *CoffeeRoaster, i *Ingredients) {
	m.espresso = *e
	m.roaster = *r
	m.ingredients = *i
}

func main() {
	Espresso := EspressoMaker{20, 0.30}
	Roaster := CoffeeRoaster{100}
	IngredientsAdder := Ingredients{}


	GordonsCoffeeMachine := new(CoffeeMaker)
	GordonsCoffeeMachine.SetMachine(&Espresso, &Roaster, &IngredientsAdder)
	GordonsCoffeeMachine.MakeCappuccino()
	GordonsCoffeeMachine.MakeLatte()
	GordonsCoffeeMachine.MakeVienna() //Coffee Beans will be out of stock by now :P
}


