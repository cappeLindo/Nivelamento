package main

import (
	"fmt"
)

func main() { 
	for dia := 1; dia <= 30; dia++ { 
		fmt.Println("\033[1mdia\033[0m: ",dia, "\n")
		for hora := 1; hora <= 24; hora++ { 
			fmt.Println("hora: ", hora, "\n")
			for min := 1; min < 60; min++ {
				fmt.Print(" ",min)
			}
			fmt.Println("\n")
		}
	}
}