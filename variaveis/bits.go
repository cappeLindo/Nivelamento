package main

import "fmt"

const (
	_ = iota 
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() { 

	fmt.Println(" ")
	fmt.Println("Binário\t\t\t\t\t\t\tDecimal")
	fmt.Println(" ")
	fmt.Printf("Bytes: %b\t\t\t\t\tDecimal: %d\n", kb, kb)
	fmt.Printf("Bytes: %b\t\t\t\tDecimal: %d\n", mb, mb)
	fmt.Printf("Bytes: %b\t\t\tDecimal: %d\n", gb, gb)
	fmt.Printf("Bytes: %b\tDecimal: %d\n", tb, tb)
	fmt.Println(" ")
	
}