package main

import "fmt"

func main() {
	a := 123456789
	b := 98765432

	k := createDivisor(a)

	for k >= 1 {
		num1 := (a / k) % 10

		l := createDivisor(b)
		for l >= 1 {
			num2 := (b / l) % 10
			if num1 == num2 {
				fmt.Printf("%v ", num1)
				break
			}
			l /= 10
		}

		k /= 10
	}
}

func createDivisor(a int) int {
	k := 1
	for k < a {
		k *= 10
	}
	k /= 10

	return k
}
