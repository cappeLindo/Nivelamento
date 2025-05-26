package main 

import "fmt"

func main() { 
	cansado := "sim"
	if cansado == "sim" { 
		fmt.Println("Vai dormir mano")
	} else if cansado == "não" { 
		fmt.Println("Vai estudar mano")
	} else { 
		fmt.Println("Faz alguma coisa mano")
	}
}

// O ex06 e 07 são praticamente o mesmo, apenas muda em que um pede-se apenas if else e no outro pede if, else e if else