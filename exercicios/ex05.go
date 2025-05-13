package main

import "fmt"

type inteiro int

var x inteiro

var y int

func main() {

	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Printf("x: %v\n",x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("y: %T", y)
}
