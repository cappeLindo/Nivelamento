package main

import (
	"fmt"
)

func main() { 
	//a := 0
	//b := 1

	//for i := 1; i <= 10; i++ { 
		//fmt.Println(b)
		//c := a
		//a = b
		//b = c + a 
	
	//}

	a := 0
	b := 1

	for i := 1; i <= 10; i++ { 
		fmt.Println(b)
		a, b = b, a+b
	}
	
}