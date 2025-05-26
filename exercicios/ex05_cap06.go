package main 

import "fmt"

func main() { 

	for i := 10; i <= 100; i++ { 
		if i % 4 == 0 { 
			fmt.Printf("%v é pode ser dividido por 4\n", i)
		}
	}
}