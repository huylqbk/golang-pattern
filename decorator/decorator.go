package main

import (
	"fmt"
)

func main() {
	fmt.Println(CostA())

	funcA := Decorator(CostA)
	funcB := Decorator(CostB)

	funcC := Decorator(funcA)

	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
}

func CostA() int {
	return 10
}

func CostB() int {
	return 15
}

func Decorator(f func() int) func() int {
	return func() int {
		fmt.Println("add func")
		return f() + 5
	}
}
