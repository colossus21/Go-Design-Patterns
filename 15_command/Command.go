package main

import "fmt"

//The Command Interface implemented by all commands

type ICommand interface {
	execute()
	unexecute()
	setReceiver(calc Device)
}

//Invoker: Bridge between Commands and Devices

type Invoker struct {
	currentIndex int
	commandQueue []ICommand
	calc Device
}

func (i *Invoker) Command(c ICommand) {
	i.commandQueue = append(i.commandQueue, c)
	i.currentIndex = len(i.commandQueue) - 1
	c.setReceiver(i.calc)
	c.execute()
}

func (i *Invoker) Undo() {
	if (i.currentIndex < 0) {
		fmt.Println("You cannot perform the undo operation anymore!!")
	} else {
		command := i.commandQueue[i.currentIndex]
		command.unexecute()
		i.commandQueue[i.currentIndex] = nil
		i.currentIndex--
	}
}

func InvokerInit(device Device) *Invoker {
	q := []ICommand {}
	i := Invoker{-1,q, device}
	return &i
}

// Calculator Interface

type Device interface {
	TurnOn()
	TurnOff()
}

// Receiver: SimpleCalculator

type SimpleCalculator struct {
}

func (s *SimpleCalculator) TurnOn() {
	fmt.Println("Calculator is turned on!")
}

func (s *SimpleCalculator) TurnOff() {
	fmt.Println("Calculator is turned off!")
}

// Receiver: TV

type TV struct {
	name string
}

func (s *TV) TurnOn() {
	if (s.name == "") {
		fmt.Println("Generic TV is turned on, channels overload !!")
	} else {
		fmt.Println(s.name, "is turned on !! Have a a good day!!")
	}
}

func (s *TV) TurnOff() {
	if (s.name == "") {
		fmt.Println("Generic TV is turned off, channels overload !!")
	} else {
		fmt.Println(s.name, "is turned off !! Have a a good day!!")
	}
}

func (s *TV) SetName(n string) {
	s.name = n
}

// Command: DeviceTurnOn

type DeviceTurnOn struct {
	calc Device
}

func (c *DeviceTurnOn) execute() {
	c.calc.TurnOn()
}

func (c *DeviceTurnOn) unexecute() {
	c.calc.TurnOff()
}

func (c *DeviceTurnOn) setReceiver(calc Device) {
	c.calc = calc
}

// Command: DeviceTurnOff

type DeviceTurnOff struct {
	calc Device
}

func (c *DeviceTurnOff) execute() {
	c.calc.TurnOff()
}

func (c *DeviceTurnOff) unexecute() {
	c.calc.TurnOn()
}

func (c *DeviceTurnOff) setReceiver(calc Device) {
	c.calc = calc
}

func main() {
	onCommand := new(DeviceTurnOn)
	offCommand := new(DeviceTurnOff)

	Calc := new(SimpleCalculator)
	CalcController := InvokerInit(Calc)
	CalcController.Command(onCommand)
	CalcController.Command(offCommand)
	CalcController.Undo()
	CalcController.Command(offCommand)
	CalcController.Undo()

	SonyTV := new(TV)
	SonyTV.SetName("Sony Bravia")
	TVController := InvokerInit(SonyTV)
	TVController.Command(onCommand)
	TVController.Command(offCommand)
	TVController.Undo()
	TVController.Undo()

	OldTV := new(TV)
	OldTVController := InvokerInit(OldTV)
	OldTVController.Command(onCommand)
	OldTVController.Undo()
	OldTVController.Undo()
}
