package main

import (
	"fmt"
)

func Varmain() {
	egInt()
}
func egInt() {
	var A1 int
	A1 = 10
	A2 := 20
	var A3 = 30
	fmt.Println(A1, A2, A3)
}
func egcont() {
	const B1 = 100
	fmt.Printf("%d , %T\n", B1, B1)
}
