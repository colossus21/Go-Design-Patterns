package main

import "fmt"

var lastUserId, lastItemId = 0, 0

type Item struct {
	id int
	name string
	ownerId int
}

//Mediator

type InventoryManger interface {
	AddItem(name string) *Item
	AddUser(name string) *User
	SetOwnerOf(itemId int, userId int)
	ShowItemsOf(userId int)
	Print()
}

type CentralInventory struct {
	items []*Item
	users []*User
}

func (i *CentralInventory) AddItem(name string) *Item {
	item := new(Item)
	lastItemId++
	item.id = lastItemId
	item.name = name
	i.items = append(i.items, item)
	return item
}

func (i *CentralInventory) AddUser(name string) *User {
	user := new(User)
	lastUserId++
	user.id = lastUserId
	user.name = name
	user.inventory = i
	i.users = append(i.users, user)
	return user
}

func (i *CentralInventory) ShowItemsOf(userId int) {
	found := false
	for n := range i.items {
		if i.items[n].ownerId == userId {
			fmt.Println("id:",i.items[n].id,"name:",i.items[n].name,"owner:",i.items[n].ownerId)
			found = true
		}
	}
	if !found {
		fmt.Println("Item(s) not found !!")
	}
}

func (i *CentralInventory) SetOwnerOf(itemId int, userId int) {
	for n := range i.items {
		if i.items[n].id == itemId {
			if i.items[n].ownerId != 0 {
				fmt.Println(i.items[n].name,"item bought by", i.GetOwnerById(userId).name, "from", i.GetOwnerById(i.items[n].ownerId).name)
			} else {
				fmt.Println(i.items[n].name,"item bought by", i.GetOwnerById(userId).name)
			}
			i.items[n].ownerId = userId
			return
		}
	}
	fmt.Println("Item not found !!")
}

func (i *CentralInventory) GetOwnerById(userId int) *User {
	for n := range i.users {
		if i.users[n].id == userId {
			return i.users[n]
		}
	}
	return nil
}

func (i *CentralInventory) Print() {
	for _, item := range i.items {
		fmt.Print("{ ")
		fmt.Print("id: ",item.id,", name: ",item.name,", owner: ",item.ownerId)
		fmt.Print(" } ")
	}
	fmt.Println()
}

//Colleague

type User struct {
	id int
	name string
	inventory InventoryManger //Mediator
}

func (u *User) BuyItem(itemId int) {
	u.inventory.SetOwnerOf(itemId, u.id)
}

func (u *User) BoughtItems() {
	fmt.Println("[ITEMS]:",u.name)
	u.inventory.ShowItemsOf(u.id)
}

func main() {
	// Mediator
	Bazaar := new(CentralInventory)
	// Add Items
	Bazaar.AddItem("Angry Doritos!")
	Bazaar.AddItem("Big Bad Cucumber!")
	Bazaar.AddItem("Fluffy Donut!")
	Bazaar.AddItem("Flying Cheetos!")
	// Print Item Info
	Bazaar.Print()
	// Add Users
	u1 := Bazaar.AddUser("Zorro")
	u2 := Bazaar.AddUser("Sasuke")
	u3 := Bazaar.AddUser("Vegeta")

	// Buy Items through Mediator
	u1.BuyItem(1)
	u2.BuyItem(2)
	u3.BuyItem(3)
	u1.BuyItem(4)

	// Prints Item(s) bought by each user
	u1.BoughtItems()
	u2.BoughtItems()
	u3.BoughtItems()

	// Buy Items to check the transaction
	u3.BuyItem(4)

	// Prints Item(s) bought by each user
	u1.BoughtItems()
	u2.BoughtItems()
	u3.BoughtItems()

	// There can be more User types or related Objects as colleagues in Mediator pattern
}