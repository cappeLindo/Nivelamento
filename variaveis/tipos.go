package main

import "fmt"

var x int = 10

type hotdog int //criamos um type que se chama hotdog e recebe como seu parâmetro que ele é int
var b hotdog = 500 // aqui definimos que b é do tipo "hotdog"

func main() { 

	fmt.Printf("%v, %T \n", x, x)
	fmt.Println(b)
	fmt.Printf("%T \n", b,)

	x = int(b)
	println(x)

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


