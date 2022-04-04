package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
	fmt.Println("workrer")
}

func main() {
	worker()
}