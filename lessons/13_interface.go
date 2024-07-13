package main

import "fmt"

type Flyer interface {
	Fly() string
}

type Duck struct {
	isAnimal bool
	name     string
	age      int
}

func (d *Duck) Fly() string {
	fmt.Println("Утка подумала")
	return "Утка полетела"
}

func (d *Duck) ChangeAge() {
	d.age += 1
}

type Docker interface {
	Flyer
}

func main() {
	var bird Flyer
	fmt.Printf("%T, %v\n", bird, bird)

	//var stupidBird *Duck
	//bird = stupidBird
	//fmt.Printf("%T, %v\n", bird, bird)

	cleverBird := &Duck{name: "utka", age: 10, isAnimal: true}
	bird = cleverBird

	bird.Fly()
	bird.Fly()

	fmt.Printf("%T, %v\n", bird, bird)
}
