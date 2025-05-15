package main 

import "fmt"

const (
	a = iota * 10
	_ 
	c 
	d 
	_ 
)
/*
const m = "Oi bom dia!"
const n = 10
const x = 10.0


var c int
var q float64
var s string
*/

func main() {
	//s = m 
	//c = n
	//q = x
	//fmt.Println(s, c, q)
	fmt.Println(a, c, d)
}