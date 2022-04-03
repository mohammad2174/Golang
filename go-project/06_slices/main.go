package main

import "fmt"

func main() {
	foo := []string{"mohammad", "reza", "khorrami"}

	fmt.Printf("foo = %v (type %T)\n", foo, foo)

	fmt.Println(len(foo))
	
	fmt.Println(foo[1])

	fmt.Println(foo[1:])

	for i := 0; i < len(foo); i++ {
		fmt.Println(foo[i])
	}

	for index := range foo {
		fmt.Println(index)
	}

	for index, value := range foo {
		fmt.Printf("%s at %d\n", value, index)
	}

	for _, value := range foo {
		fmt.Println(value)
	}

	foo = append(foo, "Tehran")
	fmt.Println(foo)
}