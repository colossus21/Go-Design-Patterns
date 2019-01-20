package main

import "fmt"

type Drinks interface {
	ServeDrink()
}

type Tea interface {
	GetTealeavesName() string
}

type Coffee interface {
	GetRoastType() string
}

//Embedded Composition

type Mocha struct {
	Drinks
	Coffee
}
func (m *Mocha) GetRoastType() string {
	return "Dark Roast"
}
func (m *Mocha) ServeDrink() {
	fmt.Println("A cup of Mocha with " + m.GetRoastType() + " type is served!!")
}
type GreenTea struct {
	Drinks
	Tea
}
func (t *GreenTea) GetTealeavesName() string {
	return "Chinese Green Aroma"
}
func (t *GreenTea) ServeDrink() {
	fmt.Println("Green tea of " + t.GetTealeavesName() + " is served!")
}
//FOR DIRECT COMPOSITION

type Drink struct {
	name string
}
func (d *Drink) ServeDrink() {
	fmt.Println(d.name + " has been served !!")
}

func main() {
	coffee := Mocha{}
	tea := GreenTea{}
	coffee.ServeDrink()
	tea.ServeDrink()

	//Direct composition example struct inside struct
	type Espresso struct {
		this_coffee Drink
		GetRoastType func() string
	}
	type BlackTea struct {
		this_tea Drink
		GetTealeavesName func() string
	}

	EspressoRoastType := func() string {
		return "Medium Roast"
	}

	BlackTeaGetTealeavesName := func() string {
		return "Chinese Black Aroma"
	}
	coffee2 := Espresso{Drink{"Espresso"},EspressoRoastType}
	tea2 := BlackTea{Drink{"Black Tea"}, BlackTeaGetTealeavesName}
	coffee2.this_coffee.ServeDrink()
	tea2.this_tea.ServeDrink()
}

