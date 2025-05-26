package main

import "fmt"

func main() {

	ano:= 2008
	for { 
		if ano != 2007 { 
			fmt.Printf("Você não nasceu em %v", ano)
			break
		} else { 
			fmt.Printf("Isso aí, você nasceu em %v", ano)
			break
		}
	}
}