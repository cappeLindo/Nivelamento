package main	

import "fmt"

func main() { 
	horaDoDia("dia")
	horaDoDia("Dia")
}

func horaDoDia(s string) string { 
	if s == "dia" || s == "Dia" { 
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" || s == "Tarde" { 
		fmt.Println("Oi, boa tarde!")
	} else if s == "noite" || s == "Noite" { 
		fmt.Println("Oi, boa noite!")
	} else {
		fmt.Println("[erro]")
	}
	return s
}