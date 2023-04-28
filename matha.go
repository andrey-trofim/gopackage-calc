package gopackage_calc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multi(a, b int) int {
	return a + b
}

func PrintNumbers(a, b int) {
	fmt.Println(a,b)
}

func Reverse(a, b int) []int {
	r := []int{b, a}
	return r
}