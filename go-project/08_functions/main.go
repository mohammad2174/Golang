package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divideModular(a int, b int)  (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divideModular(7, 2)
	fmt.Printf("div = %d, mod = %d", div, mod)
}