package main

import "fmt"

func main() { 
	
	fmt.Println("os números pares de 1 á 20 são: ")
	for i := 1; i <= 20; i ++ { 
		if i % 2 == 0 {
			fmt.Printf("Número par:   %v\n", i)
			} else { 
				fmt.Printf("Número impar: %v\n", i)
			}
	}
}