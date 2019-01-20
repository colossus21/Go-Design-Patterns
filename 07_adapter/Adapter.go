package main

import "fmt"

type Coffee struct {
	sugar int
	milk int
}

func (c *Coffee) GetInfo() {
	fmt.Println("This coffee contains - ", c.sugar, "gm sugar and ", c.milk, "ml milk.")
}

//A 21st Century Coffee Maker

type CoffeeMaker interface {
	makeCoffee(sugarAmount int, milkAmount int) Coffee
}

//Let's make a CoffeeMaker based on the 21st Standard

type ClassicCoffeeMachine struct {
	milkAmount int
	sugarAmount int
}

func (cm *ClassicCoffeeMachine) makeCoffee(sugarAmount int, milkAmount int) Coffee{
	fmt.Println()
	fmt.Println("CLASSIC MACHINE")
	fmt.Println("Sugar added: ",sugarAmount,"gm, Milk added: ",milkAmount,"ml")
	if(sugarAmount<cm.sugarAmount && milkAmount<cm.milkAmount) {
		cm.milkAmount -= milkAmount
		cm.sugarAmount -= sugarAmount
		fmt.Println("Remaining [Sugar Milk] =", cm.sugarAmount, cm.milkAmount)
	} else {
		panic("Not enough Milk and Sugar, turning off...")
	}
	return Coffee{sugar:sugarAmount,milk:milkAmount}
}

//A 22nd Century Coffee Maker doesn't require any parameters, it detects the amounts of milk and sugar based on your mood

type NewCoffeeMaker interface {
	makeCoffee() Coffee
}

//The new CofeeMaker might look like a new Machine but the underlying process is the same as the classic one
//We need to make the Coffee with the classic makeCofee() function, but for the sake of new interface we need to adapt to its needs
//This machine uses an adaptor to make things work

type NewCoffeeMachine struct {
	milkAmount int
	sugarAmount int
	classicMachine CoffeeMaker
}

func (nm *NewCoffeeMachine) makeCoffee() Coffee {
	//We assume, the machine detected the optimal amount of sugar and milk will be 100gm and 250ml
	fmt.Println()
	fmt.Println("BRAND NEW MACHINE:")
	coffee := nm.classicMachine.makeCoffee(100,250)
	fmt.Println("ADDING SOME TOUCHES OF NEWNESS!!")
	fmt.Println("AAH!! THE AROMA...")
	return coffee
}

func main() {
	OldMachine := new(ClassicCoffeeMachine)
	OldMachine.milkAmount = 500
	OldMachine.sugarAmount = 300

	ClassicCoffee := OldMachine.makeCoffee(200, 100)
	ClassicCoffee.GetInfo()

	//Let's Upgrade our OldMachine to the NewOne
	NewMachine := new(NewCoffeeMachine)
	NewMachine.milkAmount = 1000
	NewMachine.sugarAmount = 750
	NewMachine.classicMachine = OldMachine
	OldMachine.milkAmount = NewMachine.milkAmount
	OldMachine.sugarAmount = NewMachine.sugarAmount
	NewCoffee := NewMachine.makeCoffee()
	NewCoffee.GetInfo()
}