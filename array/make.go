package main

import "fmt"

func main() { 
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice)

	//agrupando ilices dentro de uma slice, isso se chama slice mult-dimensional

	ss := [][] int { 
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9, 10},
	}

	//aqui uso o len() para me retornar quantos elementos tem dentro da slice
	fmt.Println(len(ss[2]))
	
	fmt.Println(ss [1] [2])
}