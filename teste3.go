//Função dentro de função
package main

import "fmt"

var x int

func f() {
	x = 123
	g := func() {
		y := 321
		fmt.Println(y)
	}
	g()
}

func main() {
	f()
	fmt.Println(x)
}
