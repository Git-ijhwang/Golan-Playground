package main

import "fmt"

type trade struct {
	symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t := trade{
		symbol: "TSLA",
		Volume: 127,
		Price:  625,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
