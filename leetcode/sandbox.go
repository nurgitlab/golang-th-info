package main

import "fmt"

type Man struct {
	age  int
	name string
}

func (m *Man) canSay(s string) {
	fmt.Printf("My name is %s and I am %s\n", m.name, s)
}

type Troop struct {
	Man
	division string
}

func (t *Troop) canDo() {
	fmt.Println("Yes SIR!")
}

type SoldierCan interface {
	canDo()
}

func main() {
	var a SoldierCan = &Troop{
		Man: Man{
			age:  21,
			name: "Ivan",
		},
		division: "Ru division",
	}

	b := &Troop{
		Man: Man{
			age:  20,
			name: "Old Ivan",
		},
		division: "Su div",
	}

	a.canDo()
	//a.canSay()

	b.canSay("fff")
	b.canDo()
}
