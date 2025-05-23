package main

import "fmt"

func main() { 
	
	num := 6
	resultado := 1

	for cont := num; cont > 1; cont-- { 
		resultado *= cont
	}
	fmt.Printf("O fatorial de %v! é: %v",num , resultado)
}
