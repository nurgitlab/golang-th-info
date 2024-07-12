package main

import "fmt"

func main() {
	var p1 **int
	var p2 *int

	a := 10
	p2 = &a
	p1 = &p2
	
	fmt.Println(*p2)
	fmt.Println(**p1)

	f(p2)
	fmt.Printf("Это число %#v", a)
}

func f(p *int) {
	*p = 228
}

//func main() {
//	var b *int
//	a := 10
//	fmt.Println(b)
//	b = &a
//	fmt.Println(b)
//	fmt.Println(*b)
//
//	a = 20
//	fmt.Println(b)
//	fmt.Println(*b)
//}
