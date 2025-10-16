package main

import "fmt"

type Student struct {
	Name string
	Age  int
	City string
}

func main() {

	st := Student{Name: "Али", Age: 17, City: "Грозный"}

	ps := &st
	ps.Age = 18
	ps.City = "Аргун"

	fmt.Println(st)

}
