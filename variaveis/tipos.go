package main

import "fmt"

var x float64 = 10.1

func main() { 

	fmt.Printf("%v, %T \n", x, x)

	declarar()
}

var y int //aqui estou declarando

func declarar() { 

	y = 50 // aqui estou inicializando e atribuindo 
	fmt.Println("inicilizando:", y) 
	
	y = 20 // aqui estou atribuindo 
	fmt.Println("atribuindo:", y)
}



//declaração, inicialização e atribuição


