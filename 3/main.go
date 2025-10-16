package main

import "fmt"

func main() {

	var pricePtr *float64
	price := 199.0
	pricePtr = &price

	fmt.Println(pricePtr == nil, *pricePtr)
}
