package main

import "fmt"

// Not recommended
// Only one instance of the class

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func main() {
	i1 := GetInstance()
	i2 := GetInstance()
	fmt.Println("i1 - ", i1.AddOne())
	fmt.Println("i1 - ", i1.AddOne())
	fmt.Println("i2 - ", i2.AddOne())
}