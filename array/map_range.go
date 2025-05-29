package main

import "fmt"

func main() { 

	comidas := map[int]string { 
		1 : "Churrasco",
		2 : "Hamburger",
		3 : "Pizza",
		4 : "Risoto de frango",
		5 : "Hot filadélfia",
		6 : "Temaki",
	}

	fmt.Println(comidas)

	for valor, indice := range comidas { 
		fmt.Printf("%v. %v\n", valor, indice)
	}

	// agora para deletar é simples apenas coloque delete({nome do map}, {indice do map, Num||Str})
	delete(comidas, 6)
	fmt.Println(comidas)
}