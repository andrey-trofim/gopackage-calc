package gopackage_calc

import "fmt"

func Add(a, b int) int {
	return a + b
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

func Double(a, b int) []int {
	var r []int
	for i := range []int{a,b} {
		r = append(r, 2*i)
	}
	return r
}

func Just(a, b int) string {
	return "text"
}

func Check(a, b int) int {
	return 5
}

func First() string {
	return "first"
}