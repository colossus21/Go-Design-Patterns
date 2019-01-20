package main

import "fmt"

type Payment interface {
	Pay() error
}

//Payment contains two concretes in this example

type PayByCash struct {
	amount int
}

func (c *PayByCash) Pay() error {
	fmt.Println("Pay By Cash!!")
	return nil
}

type PayByCard struct {
	cardType string
	cardName string
	cardNumber int
}

func (c *PayByCard) Pay() error {
	fmt.Println("Pay By Card!!")
	return nil
}

//We need a context class to implement the strategy pattern

type Shopper struct {
	payMethod Payment
}

func (s *Shopper) setPayMethod(m Payment) {
	s.payMethod = m
}

func (s *Shopper) Pay() error {
	return s.payMethod.Pay()
}

type Admin struct {
	payMethod Payment
}

func (s *Admin) setPayMethod(m Payment) {
	s.payMethod = m
}

func (s *Admin) Pay() error {
	return s.payMethod.Pay()
}

func main() {
	Card := PayByCard{}
	Cash := PayByCash{}
	Customer := Shopper{}
	Admin := Admin{}
	

	Customer.setPayMethod(&Card)
	Customer.Pay()

	Admin.setPayMethod(&Cash)
	Admin.Pay()
}
