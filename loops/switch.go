package main

import "fmt"

func main() { 
	x := 1
	switch { 
		case x >= 1 && x <= 3:
			fmt.Printf("Estamos no Verão")
		case x >= 4 && x <= 6:
			fmt.Println("Estamos no Outono")
		case x >= 7 && x <= 9:
			fmt.Println("Estamos no Inverno")
		case x >= 10 && x <= 12:
			fmt.Println("Estamos na Primavera")
		default:
			fmt.Println("Colocar o número certo referente aos meses do ano")
		}   
	}