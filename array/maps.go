package main 

import "fmt"

func main() {

	amigos := map[string]int { 
		"Pretin" : 16,
		"Haigo" : 16,
		"Vini" : 18,
		"Guei" : 17,
	}

	fmt.Println(amigos["Pretin"])

	amigos["Carlos"] = 17
	amigos["Nathan"] = 17
	fmt.Println(amigos)

	

	if consultando_amigos, ok := amigos["Guei"]; !ok { 
		fmt.Println("não existe!")
	} else { 
		fmt.Println(consultando_amigos, "é um de seus amigos")
	}
}