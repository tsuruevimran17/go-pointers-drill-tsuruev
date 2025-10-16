package main

import "fmt"

func main() {

	temperature := 24
	tPtr := &temperature
	temperature = *tPtr + 5

	fmt.Println(temperature)
}
