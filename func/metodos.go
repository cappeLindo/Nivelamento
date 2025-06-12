package main

import "fmt"

type pessoa struct { 
	nome	string
	idade	int8
}

func (p pessoa) oibomdia() { 
	fmt.Println(p.nome, "diz bom dia")
}

func main() { 

	vitor := pessoa{
		nome: "Vitor",
		idade: 17,
	}
	vitor.oibomdia()
}

//métodos é especifico a tipos