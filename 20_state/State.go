package main

import "fmt"

// State pattern explore all possible states of the implemented model
// To implement state pattern, firstly, we find the possible states of the model and its use case actions
// State interface should contain all use case methods of the model
// Each state inherits the State interface

// It is recommended to draw a state diagram of the model before implementation of the code

// The state diagram for the following sample model:
// --YetToOrder-->(Order)--Ordered-->(CancelOrder) or,
// --YetToOrder-->(Order)--Ordered-->(Deliver)--Delivered-->(Pay)
// The initial state is [YetToOrder] and the target function is (Pay)
// Once we reach [Delivered] state and the user pays, the state is reset to [YetToOrder]

//For 3 states and 4 use case methods we get 3x4=12 possible sates of the application

type State interface {
	order()
	cancelOrder()
	deliver()
	pay()
}

type ItemOrder struct {
	itemName string
	itemState State
}

func (o *ItemOrder) order() {
	o.itemState.order()
}

func (o *ItemOrder) cancelOrder() {
	o.itemState.cancelOrder()
}

func (o *ItemOrder) deliver() {
	o.itemState.deliver()
}

func (o *ItemOrder) pay() {
	o.itemState.pay()
}

func (o *ItemOrder) setState(state string) {
	switch state {
	case "YetToOrder":
		o.itemState = &YetToOrder{o}
	case "Ordered":
		o.itemState = &Ordered{o}
	case "Delivered":
		o.itemState = &Delivered{o}
	default:
		fmt.Println("Unknown state!")
	}
}

func NewItem(name string) *ItemOrder {
	item := new(ItemOrder)
	item.itemName = name
	item.setState("YetToOrder")
	return item
}

// State: YetToOrder

type YetToOrder struct {
	itemOrder *ItemOrder
}

func (s *YetToOrder) order() {
	fmt.Println("Ordered the item", s.itemOrder.itemName)
	s.itemOrder.setState("Ordered")
}
func (s *YetToOrder) cancelOrder() {
	fmt.Println("You have no orders to cancel!")
}
func (s *YetToOrder) deliver() {
	fmt.Println("Deliver what? You haven't even placed an order.")
}
func (s *YetToOrder) pay() {
	fmt.Println("Pay for? Make sure you order something first.")
}

// State: Ordered

type Ordered struct {
	itemOrder *ItemOrder
}

func (s *Ordered) order() {
	fmt.Println("You have already ordered! You can reorder once the delivery is complete.")
}
func (s *Ordered) cancelOrder() {
	fmt.Println("Your order for item", s.itemOrder.itemName, "is cancelled!")
	s.itemOrder.setState("YetToOrder")
}
func (s *Ordered) deliver() {
	fmt.Println("Your order for item", s.itemOrder.itemName, "is delivered!")
	s.itemOrder.setState("Delivered")
}
func (s *Ordered) pay() {
	fmt.Println("Pay once we deliver the product!")
}

// State: Delivered

type Delivered struct {
	itemOrder *ItemOrder
}

func (s *Delivered) order() {
	fmt.Println("Product is delivered, please pay your due first to reorder items.")
}
func (s *Delivered) cancelOrder() {
	fmt.Println("Cannot cancel order that is delivered already! Please pay your dues.")
}
func (s *Delivered) deliver() {
	fmt.Println("Product is already delivered!")
}
func (s *Delivered) pay() {
	fmt.Println("Payment successful! You can reorder items now!")
	s.itemOrder.setState("YetToOrder")
}

func main() {
	item1 := NewItem("Mexican Hot Pizza!")

	// Let's execute the methods properly for now
	item1.order()
	item1.deliver()
	item1.pay()

	item2 := NewItem("Cheese Burger!")

	// Proper Commands
	item2.order()
	item2.cancelOrder()

	// Improperly sequenced commands
	item2.order()
	item2.pay()
	item2.deliver()
	item2.cancelOrder()
	item2.pay()
}