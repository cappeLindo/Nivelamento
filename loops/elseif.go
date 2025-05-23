package main 

import "fmt"

func main() { 
	x := 10
	y := 5
	 if x > y {
		fmt.Printf("%v é maior que %v", x, y)
	 } else if x < y { 
		fmt.Printf("%v é menor que %v", x, y)
	 } else if x == y { 
		fmt.Printf("%v é igual a %v", x, y)
	 } else if x != y { 
		fmt.Printf("%v é diferente a %v", x, y)
	 } else { 
		fmt.Println("X n é nada")
	 }
}