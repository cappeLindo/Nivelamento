package main

import "fmt"

func main() { 
	
	num := 9
	fatoria := 1

	for cont := num; cont >= 1; cont-- { 
		fatoria = fatoria * cont
	}
	fmt.Println(fatoria)
}
