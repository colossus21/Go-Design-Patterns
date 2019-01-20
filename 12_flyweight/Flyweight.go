package main

import (
	"fmt"
	"time"
)

type IFlyweightFactory interface {
	getFact(n int) *int
}

type FlyweightFactory struct {
	outputs map[int]*int
}

func (f *FlyweightFactory) getOutput(n int) *int {
	if (f.outputs[n]!=nil) {
		//fmt.Println(*f.factorials[n])
		return f.outputs[n]
	} else {
		//This can optimized further
		tmp := getOutput(n)
		f.outputs[n] = &tmp
		return &tmp
	}
}

func getOutput(n int) int {
	if n>9999 {
		return n
	} else {
		sum := 0
		for i := n; i >= 0; i-- {
			sum += i*i
			for j:=i; j>= 0; j-- {
				sum -= 1
			}
		}
		return getOutput(sum)
	}
}

func main() {
	store := make(map[int]*int)
	factory := FlyweightFactory{store}
	fmt.Println(*factory.getOutput(10))

	fmt.Println("WITH FLYWEIGHT PATTERN- Outputs:")
	fmt.Println()
	fstart := time.Now()

	fmt.Println(*factory.getOutput(21))
	fmt.Println(*factory.getOutput(10))
	fmt.Println(*factory.getOutput(21))
	fmt.Println(*factory.getOutput(10))
	fmt.Println(*factory.getOutput(31))
	fmt.Println(*factory.getOutput(31))
	fmt.Println(*factory.getOutput(39))
	fmt.Println()
	fend := time.Since(fstart).Nanoseconds()
	fmt.Println("Time Taken: ", fend)



	fmt.Println("WITHOUT FLYWEIGHT PATTERN- Outputs:")
	start := time.Now()
	fmt.Println()
	fmt.Println(getOutput(21))
	fmt.Println(getOutput(10))
	fmt.Println(getOutput(21))
	fmt.Println(getOutput(10))
	fmt.Println(getOutput(31))
	fmt.Println(getOutput(31))
	fmt.Println(getOutput(39))
	fmt.Println()
	end := time.Since(start).Nanoseconds()
	fmt.Println("Time Taken: ", end)

	fmt.Println("Efficiency: ", (float32(end)-float32(fend))/float32(end)*100.0, "%")
}