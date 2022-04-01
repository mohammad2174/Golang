package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}