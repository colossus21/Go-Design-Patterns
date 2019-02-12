package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//Observer pattern: One or Many Observers receive updates when an Observable changes
//The Observers requires no knowledge of the Observables


//Observable/Publisher

type Publisher interface {
	register(O Observer)
	unregister(O Observer)
	notify()
}


//Observer

type Observer interface {
	update(n int)
}

//Random Number Generator Observable

type RandomNumberGenerator struct {
	observers []Observer
	delay time.Duration
	rand int
}

func (g *RandomNumberGenerator) register(O Observer) {
	g.observers = append(g.observers, O)
}

func (g *RandomNumberGenerator) unregister(O Observer) {
	index := -1
	for i, v := range g.observers {
		if (v == O) {
			index = i
		}
	}
	if index != -1 {
		g.observers = append(g.observers[:index], g.observers[index+1:]...)
	} else {
		log.Fatal("Cannot register the observer!")
	}
}

func (g *RandomNumberGenerator) notify() {
	for _, v := range g.observers {
		v.update(g.rand)
	}
}

func (g *RandomNumberGenerator) generateNumber() {
	if g.delay == 0 {
		g.delay = 30000000
	}
	for {
		rand.Seed(time.Now().UnixNano())
		g.rand = rand.Intn(9000000000)
		fmt.Println("Generated number (from Generator):",g.rand)
		g.notify()
		time.Sleep(g.delay)
	}
}

//Wealth Describer

type WealthDescriber struct {
	num int
	generator Publisher
}

func (d *WealthDescriber) update(n int) {
	fmt.Print("(From Wealth Describer):")
	d.num = n
	if d.num > 8000000000 {
		fmt.Println("Billionaire!!")
	} else if d.num > 7000000000 {
		fmt.Println("Millionaire!!")
	} else if d.num > 6000000000 {
		fmt.Println("Pretty Rich!!")
	} else if d.num > 5000000000 {
		fmt.Println("Wealthy!!")
	} else if d.num > 4000000000 {
		fmt.Println("Unemployed!!")
	} else {
		fmt.Println("Need a home!!")
	}
}

func (d *WealthDescriber) setGenerator(p Publisher) {
	d.generator = p
	d.generator.register(d)
}

//Number of Digits Observer

type NumberOfDigitsObserver struct {
	num int
	generator Publisher
}

func (d *NumberOfDigitsObserver) update(n int) {
	fmt.Print("(From Number of Digits Observer):")
	d.num = n
	count := 0
	for d.num != 0 {

		d.num /= 10
		count = count + 1
	}
	fmt.Println(count)
}

func (d *NumberOfDigitsObserver) setGenerator(p Publisher) {
	d.generator = p
	d.generator.register(d)
}

func main() {
	GeneratorObservable := new(RandomNumberGenerator)
	GeneratorObservable.delay = 3000000000

	WealthObservable := WealthDescriber{}
	WealthObservable.setGenerator(GeneratorObservable)

	NumberOfDigitsObservable := NumberOfDigitsObserver{}
	NumberOfDigitsObservable.setGenerator(GeneratorObservable)

	//Observables perform actions which are handled by Observers' update method
	GeneratorObservable.generateNumber()
}