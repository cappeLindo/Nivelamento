package main

import "fmt"

func main() { 
	
	fmt.Println("os números pares de 1 á 20 são: ")
	for i := 1; i <= 20; i ++ { 
		if i % 2 == 0 {
			fmt.Println(i)
		} 
	}
}