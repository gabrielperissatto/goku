//Testes Interpretador Go:

//1. Print Comum
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

//2. Print sem '}'
package main

import "fmt"

func main() {
	fmt.Println("Hello World")

//3. Declaração e atribuição de variável
package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Println(a)
}

//4. Declaração curta de variável
package main

import "fmt"

func main() {
	b := 5
	fmt.Println(b)
}

//5. Definição de função
package main

import "fmt"

func main() {
	fmt.Println(soma(1, 2))
}

func soma(x, y int) int {
	return x + y
}
