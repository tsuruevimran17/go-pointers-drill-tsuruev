package main

import "fmt"

func main(){
	a := 11
	b := 11
	fmt.Println(&a == &b,a == b) // вывод разный потому что в первом случае сравниваем адреса 
	                             // а во втором значения
}