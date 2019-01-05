package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type OrderTime string

type Item struct {
	name string
	description string
	otherInfo string
}

type Order struct {
	items []Item
	orderTime *OrderTime
	tableName string
}


var presetItems = []Item{
	{"Cappuccino", "1 Shot of espresso, Steamed Milk, Milk Foam, Sprinkled Chocolate", "Diabetes Sugar"},
	{"Mocha", "1 Shot of espresso, Steamed Milk, Milk Foam, Sprinkled Chocolate", "Caramel on Top"},
	{"Espresso", "1 Shot of espresso in an espresso cup", "No Sugar"},
}

//Types of Proxy: Remote Proxy, Virtual Proxy, Protection Proxy
//Ie. Cache Implementation, Restrict Function Calls, Dynamic Remote Handling

type IOrders interface {
	Order(items []Item, orderTime *OrderTime, tableName string) error
}

type Orders []Order
const lineLength = 60
func dashedLine() string {
	str := strings.Repeat("=",lineLength)
	fmt.Println(str)
	return str
}
func (o *Orders) Order(items []Item, orderTime *OrderTime, tableName string) error {
	dashedLine()
	//Return error if some requirement fails; not covered here. ( Should've used it in all methods :P )
	fmt.Println("Oder placed for table : ", tableName, ", Time of Order: ", *orderTime + " (24-Hour)")
	dashedLine()
	fmt.Println("Items: ")
	for i,v := range items {
		fmt.Println(i, v.name)
	}
	dashedLine()
	newOrder := Order{items,orderTime,tableName}
	*o = append(*o, newOrder)
	return nil
}

//According to Proxy pattern it has to implement the IOrders interface

type OrdersProxy struct {
	orders *Orders //We are taking the value of the Orders we defined in line 33 rather than defining the type
	noOfOrders int
	orderLimit int
}

func (o *OrdersProxy) Order(items []Item, orderTime *OrderTime, tableName string) error {
	if (o.noOfOrders>=o.orderLimit) {
		return errors.New("All cooks are working on the previous orders/ Please try again after sometimes!")
	}
	err := o.orders.Order(items, orderTime, tableName)
	if (err != nil) {
		return errors.New("Something went wrong. Please check if the values you have provided are correct!!")
	}
	o.noOfOrders++
	return nil
}

func main() {
	GordonsKitchen := new(OrdersProxy)
	RamseyOrders := new(Orders)

	GordonsKitchen.orders = RamseyOrders
	GordonsKitchen.orderLimit = 5

	currentTime := time.Now()
	hour := currentTime.Hour()
	mins := currentTime.Minute()
	currentTimeString := OrderTime(strconv.Itoa(hour) + ":" + strconv.Itoa(mins))

	GordonsKitchen.Order(presetItems, &currentTimeString, "Regular69")
	GordonsKitchen.Order([]Item {
		{"Double Latte", "4x Espresso + Lots of Milk and Cream", "Ice"},
		{"Chocolate Fountain", "Latte with Ice-cream", "Extra Sugar"},
	}, &currentTimeString, "VIP420")


}