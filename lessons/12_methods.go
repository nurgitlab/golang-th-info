package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) changeAge(newAge int) {
	s.age = newAge
}

func main() {
	s := Student{
		name: "Nurbek",
		age:  22,
	}
	s.changeAge(32)
	//Go автоматически по входному типу определяет, указатель или значение мы вызвали
	fmt.Println(s)
}
