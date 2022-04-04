package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price float64
	Buy bool
}

func main() {
	t1 := Trade{"APPLE", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{
		Price: 99.97,
		Symbol: "MICROSOFT",
		Volume: 15,
		Buy: false,
	}

	fmt.Println(t2)
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Println(t3)
	fmt.Printf("%+v\n", t3)

}