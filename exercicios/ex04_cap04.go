package main

import "fmt"

var x int
var z int

func main() {

	x = 429
	fmt.Printf("decimal: %v,\nbinário: %b, \nhexadecimal: %#x\n", x, x, x)
	fmt.Println("  ")
	z = x << 1
	fmt.Printf("decimal: %v,\nbinário: %b, \nhexadecimal: %#x", z, z, z)
	
}