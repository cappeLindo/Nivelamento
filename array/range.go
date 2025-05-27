package main

import (
	"fmt"
)

func main() { 
	slice := []string{"banana", "maça", "melancia", "morango", "laranja"}

	slice = append(slice, "Abacate")

	for indice, valor_i := range slice { 
		fmt.Println("No indice", indice, "temos o valor:",valor_i)
	}
	fmt.Println("---------------------------")

	for indice, _ := range slice { 
		fmt.Println("Esse slice tem", indice, "elementos!")
	}
	fmt.Println("---------------------------")
	fmt.Println("esse slice tem", len(slice) , "elementos!")
	fmt.Println("---------------------------")
	
	for cont, valor := range slice {
		fmt.Printf("%v. Um dos elementos dentro do slice é '%v'.\n",cont + 1, valor)
	}
	fmt.Println("---------------------------")

	slice2 := []int{20, 30, 40, 10}

	total := 0

	for _, valor := range slice2 {
		total += valor
	}
	fmt.Println("o Valor total desse slice é de:", total)
}  