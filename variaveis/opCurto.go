package main

import "fmt"

func main() { 

	x := 10
	y := "Olá bom dia"
	
	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n", y, y)

	x, z := 20, 30

	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("z: %v, %T \n", z, z)
	fmt.Printf("x: \n %v \n", x)
	fmt.Printf("x: \n %T \n", x)

	a := 10 + 15
	b := 10 == 5
	c := 10 > 1
	fmt.Println(a, "\n", b, "\n", c)
}

//:= já adiciona automáticamente o tipo da variável, e só é usada dentro de um codeBlock ou mais conhecida com func ou função
// %v exibe o valor da variável como por exemplo "x = 20, fmt.Printf("x: %v") -> x: 20" 
// %T exibe o tipo da variável como exemplo "fmt.Printf("x: %T") -> x: int"
