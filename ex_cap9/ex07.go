package main

import "fmt"


func main() { 
	pessoas := [][]string { 
		[]string{
			"Carlos",
			"Almeida",
			"Jogar",
		}, 
		[]string { 
			"Julia",
			"Alberk",
			"Estudar biologia",
		},
		[]string { 
			"João",
			"Pedro",
			"Dirigir",
		},
	}
	for _, indice := range pessoas { 
		fmt.Println(indice[0])
		for _, item := range indice { 
			fmt.Printf("\t%v\n", item)
		}
	}
}