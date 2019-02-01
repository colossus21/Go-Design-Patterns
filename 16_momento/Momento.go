package main

import "fmt"

//Memento

type Memento struct {
	state string
}

func (m *Memento) getState() string {
	return m.state
}

func (m *Memento) setState(s string) {
	m.state = s
}

type Originator struct {
	state string
}

//Originator

func (o *Originator) getState() string {
	return o.state
}

func (o *Originator) setState(s string) {
	o.state = s
}

func (o *Originator) newMemento() Memento {
	return Memento{o.state}
}

func (o *Originator) restoreFromMemento(m Memento) {
	o.state = m.getState()
}

type Caretaker struct {
	savedStates []Memento
}

//Caretaker

func (c *Caretaker) addMemento(m Memento) {
	fmt.Println("Saved state:",m.getState())
	c.savedStates = append(c.savedStates, m)
}

func (c *Caretaker) getMemento(i int) Memento {
	if i>=0 && i<len(c.savedStates) {
		return c.savedStates[i]
	} else {
		return Memento{}
	}
}

func (c *Caretaker) undo() string{
	i := len(c.savedStates) - 1
	c.savedStates = c.savedStates[0:i]

	lastSavedState := c.getMemento(i-1).state
	fmt.Println("Reverting to last saved state:", lastSavedState)
	return lastSavedState
}

func main()  {
	O := Originator{}
	C := Caretaker{}

	O.setState("Walkking!")
	O.setState("Walking!")
	C.addMemento(O.newMemento())

	O.setState("Running!")
	C.addMemento(O.newMemento())

	O.setState("Swimming!")
	C.addMemento(O.newMemento())

	//Undo State
	O.setState(C.undo())
	O.setState(C.undo())
	fmt.Println("Current state: ", O.getState())
}