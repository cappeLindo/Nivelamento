package main 

import "fmt"

func main() { 
	
	x := 10
	switch {
	case x >= 1 && x <= 3:
		fmt.Println("Você não está cansado, apenas continue")
	case x >= 4 && x <= 6: 
		fmt.Println("Tá cansadin, descansar um pouco")
	case x >= 7 && x <= 9: 
		fmt.Println("Só vai descansar")
	case x == 10:
		fmt.Println("Tira uma folga")
	default:
		fmt.Println("coloque o seu nível de cansaço")
	}
}