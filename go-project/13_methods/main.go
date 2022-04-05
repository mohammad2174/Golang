package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price float64
	Buy bool
}

func (trade *Trade) value() float64 {
	value := float64(trade.Volume) * trade.Price
	if trade.Buy {
		value = -value
	}
	return value
}

func main() {
	t := Trade{"APPLE", 10, 99.98, true}
	fmt.Println(t.value())
}