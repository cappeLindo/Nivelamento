package main 

import "fmt" 

func main() { 

	for x := 65; x <= 90; x++ { 
		fmt.Printf("%U '%v'\n", x, string(x))
	}
}