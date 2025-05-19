package main

import "fmt"

const (
	_ = 2000 + iota
	a 
	b
	c 
	d
)

func main() { 

	fmt.Printf("ano: %v\nano: %v\nano: %v\nano: %v", a, b, c, d)

}