package main

import (
	"fmt"
)

//Terms to remember: Open-Close Principle: STRUCTURES are open to EXTENSION but Closed for MODIFICATION
//IPizza will be the implemented by the all Pizzas and Decorators

type IPizza interface {
	add(pizza IPizza) IPizza
	getDescription() string
	getCost() float32
}

//Concrete classes

//Neapolitan Pizza

type Neapolitan struct {
	crust string
}
func (p *Neapolitan) add(pizza IPizza) IPizza {
	return p
}
func (p *Neapolitan) getDescription() string {
	desc := "Neapolitan Pizza"
	return desc
}
func (p *Neapolitan) getCost() float32 {
	return 500
}

//Sicilian Pizza

type Sicilian struct {
	crust string
}
func (p *Sicilian) add(pizza IPizza) IPizza {
	return p
}
func (p *Sicilian) getDescription() string {
	desc := "Sicilian Pizza"
	return desc
}
func (p *Sicilian) getCost() float32 {
	return 800
}

//Abstract Decorator

type PizzaDecorator struct {
	pizza IPizza
	price float32
}
func (d *PizzaDecorator) setPizza(pizza IPizza) IPizza {
	fmt.Println("Selected Pizza: ", pizza.getDescription(), " Price: ", pizza.getCost())
	d.price += pizza.getCost()
	return pizza
}
func (d *PizzaDecorator) add(pizza IPizza) IPizza {
	d.pizza = pizza
	d.price += d.pizza.getCost()
	fmt.Println("Added: " + d.pizza.getDescription(), " Price: ", pizza.getCost())
	return d.pizza
}

func (d *PizzaDecorator) getDescription() string {
	return "Pizza Decorator"
}

func (d *PizzaDecorator) getCost() float32 {
	return 0
}

//Cheese

type Cheese struct {
	pizza IPizza
	price float32
}

func (d *Cheese) add(pizza IPizza) IPizza {
	d.pizza = pizza
	d.price += d.pizza.getCost()
	fmt.Println("Added: " + d.pizza.getDescription())
	return d.pizza
}

func (d *Cheese) getDescription() string {
	return "Cheese"
}
func (d *Cheese) getCost() float32 {
	return 120
}



//BBQSauce

type BBQSauce struct {
	pizza IPizza
	price float32
}

func (d *BBQSauce) add(pizza IPizza) IPizza {
	d.pizza = pizza
	d.price += d.pizza.getCost()
	fmt.Println("Added: " + d.pizza.getDescription())
	return d.pizza
}

func (d *BBQSauce) getDescription() string {
	return "BBQSauce"
}
func (d *BBQSauce) getCost() float32 {
	return 60
}


func main() {
	Pizza := new(PizzaDecorator)
	Pizza.setPizza(new(Sicilian))
	Pizza.add(new(BBQSauce))
	Pizza.add(new(Cheese))
	fmt.Println("Total Price: ", Pizza.price)
}