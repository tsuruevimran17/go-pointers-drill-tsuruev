package main

import "fmt"

func inc(n *int) {
    fmt.Println(*n + 1)
}

func incCopy(n int){
	fmt.Println(n)
}

func main(){
	x:=5
	inc(&x)
	incCopy(x)
}