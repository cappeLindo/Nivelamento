package main

import "fmt"

type veiculo struct{
	portas		int8
	cor 		string
} 

type caminhonete struct{
	veiculo 		
	tracao 	bool
}

type sedan struct{
	veiculo
	luxo 	bool
}

func main() { 
	veiculo1 := veiculo{
		portas: 2,
		cor: "Cinza",
	}

	veiculo2 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor: "Preto fosco",
		},
		tracao: true,

	}
	fmt.Println(veiculo1, veiculo2)
}