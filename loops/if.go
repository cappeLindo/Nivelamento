package main

import "fmt"

func main() {
	y := 20
	if x := 100; x > y { 
		fmt.Printf("%v é maior que %v!", x, y)
		} else if x < 100 {
			fmt.Printf("%v é menor que %v!", x, y)
	}
}