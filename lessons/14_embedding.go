package main

import "fmt"

// embedding - встраивание

type Person struct {
	name string
	age  int
}

type Builder struct {
	Person
}

func (p *Person) printName() {
	fmt.Println("name:", p.name)
}

func main() {
	b := Builder{Person{name: "Vasya", age: 10}}
	fmt.Println(b)
	fmt.Println(b.Person.name)
	fmt.Println(b.age) //сработает потому что синтаксический сахар, рекурсивно проходится go
	b.printName()
}
