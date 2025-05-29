package main

import "fmt"

func main() { 
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice2 := []int{6,7,8,9,10}
	fmt.Println(slice2, 10, 11, 12, 13)
	slice = append(slice,  slice2...)
	fmt.Println(slice)
}