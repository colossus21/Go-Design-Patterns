package main

import (
	"fmt"
)

//Terms to remember: Open-Close Principle: STRUCTURES are open to EXTENSION but Closed for MODIFICATION
//IAdder will be the implemented by the Decorator and all Decoration classes

type IAdder interface {
	add(coffee IAdder) IAdder
	getDescription() IAdder
	getPrice() int
}

//Initial Struct

//Decorator
type Coffee struct {
	price int
}

func (c *Coffee) getDescription() IAdder {
	fmt.Println("Base Coffee")
	return c
}

func (c *Coffee) add(coffee IAdder) IAdder {
	coffee.getDescription()
	c.price += coffee.getPrice()
	fmt.Println("Current Price: ", c.price)
	return c
}

func (c *Coffee) getPrice() int {
	return c.price
}

func NewCoffee() IAdder {
	newCoffee := new(Coffee)
	newCoffee.price = 500
	newCoffee.getDescription()
	fmt.Println("Current Price: ",newCoffee.price)
	return newCoffee
}

//Decorations

//Syrup

type Syrup struct {}

func (c *Syrup) getDescription() IAdder {
	fmt.Println("Chocolate Syrup!!")
	return c
}

func (c *Syrup) getPrice() int {
	return 100
}

func (c *Syrup) add(coffee IAdder) IAdder {
	coffee.getDescription()
	return c
}

//Syrup

type CaramelTop struct {}

func (c *CaramelTop) getDescription() IAdder {
	fmt.Println("Caramel on Top!!")
	return c
}

func (c *CaramelTop) add(coffee IAdder) IAdder {
	coffee.getDescription()
	return c
}

func (c *CaramelTop) getPrice() int {
	return 200
}

//Promo

type Promo struct {
	prevPrice int
	percent int
}

func (c *Promo) getDescription() IAdder {
	fmt.Println("Promo added: ", c.percent, "%")
	return c
}

func (c *Promo) add(coffee IAdder) IAdder {
	coffee.getDescription()
	return c
}

func (c *Promo) getPrice() int {
	var dec float32 = (float32(c.percent)/100)
	fmt.Println(dec*float32(c.prevPrice))
	return int(dec*float32(c.prevPrice))*-1
}

func NewPromo(price int, percent int) *Promo {
	newPromo := Promo{price, percent}
	return &newPromo
}

func main() {
	c := NewCoffee()
	c.add(new(CaramelTop)).add(new(Syrup)).add(NewPromo(c.getPrice(), 15))
}


