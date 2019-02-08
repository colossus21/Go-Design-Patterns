package main

import (
	"fmt"
)

//Interfaces: Visitor, Visitable

//Visitor

type Visitor interface {
	 visitPCGame(g PCGame) float32
	 visitPSGame(g PSGame) float32
	 visitXBoxGame(g XBoxGame) float32
}

type BlackFridayVisitor struct {
	Visitor
}

func (v *BlackFridayVisitor) visitPCGame(g *PCGame) float32 {
	//50% Off for any PC Game
	return g.price * 0.5
}

func (v *BlackFridayVisitor) visitPSGame(g *PSGame) float32 {
	//30% Off for any PlayStation Game
	return g.price * 0.7
}

func (v *BlackFridayVisitor) visitXBoxGame(g *XBoxGame) float32 {
	//20% Off for any XBox Game
	return g.price * 0.8
}

type ConsoleWeekVisitor struct {
	Visitor
}

func (v *ConsoleWeekVisitor) visitPCGame(g *PCGame) float32 {
	//No Discounts for PC Games
	return g.price * 1
}

func (v *ConsoleWeekVisitor) visitPSGame(g *PSGame) float32 {
	//50% Off for any PlayStation Game
	return g.price * 0.5
}

func (v *ConsoleWeekVisitor) visitXBoxGame(g *XBoxGame) float32 {
	//50% Off for any XBox Game
	return g.price * 0.5
}

//Visitable

type Visitable interface {
	Accept(v Visitor)
}

type PCGame struct {
	name string
	minimumSpecs string
	price float32
}

func (g *PCGame) Set(name string, specs string, price float32) {
	g.name = name
	g.minimumSpecs = specs
	g.price = price
}

func (g *PCGame) Accept(v Visitor) {
	v.visitPCGame(*g)
}

type PSGame struct {
	name string
	version string
	price float32
}

func (g *PSGame) Set(name string, version string, price float32) {
	g.name = name
	g.version = version
	g.price = price
}

func (g *PSGame) Accept(v Visitor) {
	v.visitPSGame(*g)
}

type XBoxGame struct {
	name string
	models string
	price float32
}

func (g *XBoxGame) Set(name string, modelsSupported string, price float32) {
	g.name = name
	g.models = modelsSupported
	g.price = price
}

func (g *XBoxGame) Accept(v Visitor) {
	v.visitXBoxGame(*g)
}
func main() {

	BlackFridayPrice := BlackFridayVisitor{}
	ConsoleWeekPrice := ConsoleWeekVisitor{}

	MetalGearSolid := new(PCGame)
	MetalGearSolid.Set("Metal Gear Solid", "RAM 512MB GPU 2GB", 1000.0)

	MetalGearSolidPS := new(PSGame)
	MetalGearSolidPS.Set("Metal Gear Solid", "PSv4", 1000.0)

	MetalGearSolidXBox := new(XBoxGame)
	MetalGearSolidXBox.Set("Metal Gear Solid", "XBOX 360, XBOX One", 1000.0)

	fmt.Println("Price of all versions of Metal Gear Solid Games: 1000")

	fmt.Println("Black Friday Offer for PC:", MetalGearSolid.name, BlackFridayPrice.visitPCGame(MetalGearSolid))
	fmt.Println("Black Friday Offer for XBox:", MetalGearSolidXBox.name, BlackFridayPrice.visitXBoxGame(MetalGearSolidXBox))
	fmt.Println("Black Friday Offer for PS:", MetalGearSolidPS.name, BlackFridayPrice.visitPSGame(MetalGearSolidPS))

	fmt.Println("Console Week Offer for PC:", MetalGearSolid.name, ConsoleWeekPrice.visitPCGame(MetalGearSolid))
	fmt.Println("Console Week Offer for XBox:", MetalGearSolidXBox.name, ConsoleWeekPrice.visitXBoxGame(MetalGearSolidXBox))
	fmt.Println("Console Week Offer for PS:", MetalGearSolidPS.name, ConsoleWeekPrice.visitPSGame(MetalGearSolidPS))

}
