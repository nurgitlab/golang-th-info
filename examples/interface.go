package main

import "fmt"

type Animal struct {
	age  int
	name string
	say  string
}

type Person struct {
	name string
}

type AnimalCan interface {
	Run(s string) string //будут только те методыm
}

func (a *Animal) Run(s string) string {
	fmt.Println("Животное побежало ", s)
	a.say = s
	return a.say
}

func (p *Person) Run(s string) string {
	fmt.Println("Человек пошел ", s)

	return s
}

func poli(myInterface AnimalCan, to string) {
	myInterface.Run(to)
}
func (a *Animal) Codding() {
	fmt.Println()
	fmt.Printf("Животное не программирует")
}

func main() {

	var a AnimalCan = &Animal{age: 10, say: "Nothing", name: "Бобик"}
	a.Run("rrr")
	a.Run("Мууууууу")

	poli(a, "на радугу")
	var p AnimalCan = &Person{name: "Nurbek"}
	poli(p, "на повышение")

	if human, isOk := p.(*Animal); isOk {
		fmt.Println("Это животное", human.name)
	} else {
		fmt.Println("Это не животное")
	}
}
