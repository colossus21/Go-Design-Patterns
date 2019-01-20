package main

import "fmt"

//In Chain of Responsibility, the Program tries to solve a problem through the first Chain
//If the first Chain fails, it pass it to the next one and on and on

type Object struct {
	objType string
	params map[string]int
	startChain IChain
}

func (d *Object) setChain(c IChain) {
	d.startChain = c
}

func (d *Object) describeObject() {
	d.startChain.describeObject(*d)
}



//Interface for all Chains

type IChain interface {
	setNext(c IChain) IChain
	describeObject(request Object)
}

//Coffee Describer (Chain 1)

type CoffeeDescriber struct {
	nextChain IChain
}

func (d *CoffeeDescriber) setNext(c IChain) IChain {
	d.nextChain = c
	return d
}

func (d *CoffeeDescriber) describeObject(request Object) {
	if val, ok := request.params["sugar"]; (request.objType == "Coffee" && ok) {
		switch val {
		case 50:
			fmt.Println("It's an Espresso!")
		case 100:
			fmt.Println("Must be Americano!")
		case 200:
			fmt.Println("Mocha ftw!!")
		default:
			fmt.Println("Just a cup of Regular Coffee!!")
		}
	} else {
		d.nextChain.describeObject(request)
	}
}

//Pizza Describer

type PizzaDescriber struct {
	nextChain IChain
}

func (d *PizzaDescriber) setNext(c IChain) IChain {
	d.nextChain = c
	return d
}

func (d *PizzaDescriber) describeObject(request Object) {
	if val, ok := request.params["size"]; (request.objType == "Pizza" && ok) {
		switch val {
		case 12:
			fmt.Println("Neapolitan Pizza!")
		case 10:
			fmt.Println("Sicilian Pizza!")
		case 6:
			fmt.Println("Pan Pizza!")
		default:
			fmt.Println("Normal Pizza!")
		}
	} else {
		d.nextChain.describeObject(request)
	}
}

//Garbage Describer

type GarbageDescriber struct {
	nextChain IChain
}

func (d *GarbageDescriber) setNext(c IChain) IChain {
	d.nextChain = c
	return d
}

func (d *GarbageDescriber) describeObject(request Object) {
	fmt.Println("Can't recognize", request.objType, "with params:", request.params)
}

func main()  {
	CoffeeInfo := CoffeeDescriber{}
	PizzaInfo := PizzaDescriber{}
	OtherInfo := GarbageDescriber{}
	CoffeeInfo.setNext(&PizzaInfo)
	PizzaInfo.setNext(&OtherInfo)

	TestObj1 := Object{"Coffee", map[string]int{"sugar":100}, &CoffeeInfo}
	TestObj1.describeObject()

	TestObj2 := Object{"Pizza", map[string]int{"size":10}, &CoffeeInfo}
	TestObj2.describeObject()

	TestObj3 := Object{"Fruit", map[string]int{"sugar":100, "size": 200,}, &CoffeeInfo}
	TestObj3.describeObject()


}