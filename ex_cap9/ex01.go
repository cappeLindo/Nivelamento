package main 

import "fmt"

var a[5] int

func main() { 
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	
	fmt.Println(a)

	for indice, valor := range a { 
		fmt.Printf("%v. %v, Type: %T\n", indice, valor, a)
	}


}