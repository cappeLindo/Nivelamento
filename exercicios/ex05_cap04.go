package main

import "fmt"

func main() {

	x := "Isso é uma string normal"
	y := `Isso 
	
		é uma               raw
		String line`

	fmt.Printf("String normal: %v\nRaw String Line: %v", x, y)

}

//essa string é uma raw string line, a formatação não se aplica a ela, ela é implementada na tela do jeito que o programador a deixou 