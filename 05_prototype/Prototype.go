package main

import "fmt"

//The main idea behind Prototype pattern is that, prototypes of an object is declared beforehand
//Instead of instantiating the object, we clone it to save more time as more time is required in the creation process

type Coffee struct {
	name string
	ingredients string
	roast string
}

//Let's create an interface to describe the Coffee and implement it

type IMutations interface {
	GetInfo() string
	SetInfo(info infos)
}

func (c *Coffee) GetInfo() string {
	return c.name + " contains " + c.ingredients + " and has a " + c.roast + " roast type!"
}

type infos struct {
	name, ingredients, roast string
}

func (c *Coffee) SetInfo(info infos) {
	c.name = info.name
	c.ingredients = info.ingredients
	c.roast = info.roast
}

//Let's prototype some Coffees

var mocha *Coffee = &Coffee{"Mocha", "1 Shot of espresso, 1 Spoon of Chocolate Powder, Steamed Milk, Sprinkled Chocolate", "Dark"}
var cappuccino *Coffee = &Coffee{"Cappuccino", "1 Shot of espresso, Steamed Milk, Milk Foam, Sprinkled Chocolate", "Medium"}
var espresso *Coffee = &Coffee{"Espresso", "1 Shot of espresso in an espresso cup", "Light"}


const (
	Espresso = 1
	Mocha = 2
	Cappuccino =3
)

//We need an interface to Clone the Coffee

type ICloner interface {
	Clone(n int) IMutations
}

type CoffeeMachine struct {}

func(c *CoffeeMachine) Clone(n int) IMutations {
	switch n {
	case Espresso:
		cup := *espresso
		return &cup
	case Mocha:
		cup := *mocha
		return &cup
	case Cappuccino:
		cup := *cappuccino
		return &cup
	default:
		panic("Invalid Coffee Type!")
	}
}

//Returns a CoffeeMachine ie. Cloner

func GetCoffeeMachine() *CoffeeMachine {
	return new(CoffeeMachine)
}

func main() {
	Machine := GetCoffeeMachine()
	customEspresso := Machine.Clone(Espresso)
	fmt.Println(customEspresso.GetInfo())
	fmt.Println("Addresses# ESPRESSO: ",&espresso, "  CUSTOM ESPRESSO", &customEspresso)
	customEspresso.SetInfo(infos{"Custom Espresso", "Espresso + Milk", "Medium"})
	fmt.Println(customEspresso.GetInfo())
	fmt.Println("Addresses# ESPRESSO: ",&espresso, "  CUSTOM ESPRESSO", &customEspresso)
}


