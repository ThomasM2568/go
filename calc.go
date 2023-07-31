package main

import "fmt"

func main() {
	var a int = 4
	var b int = 2
	var c float32 = 2.1

	fmt.Println("a + b  = ", a+b) // addition de la variable a et b
	fmt.Println("a - b  = ", a-b) // soustraction de la variable a et b
	fmt.Println("a * b  = ", a*b) // multiplication de la variable a et b
	fmt.Println("a / b  = ", a/b) // division de la variable a et b
	fmt.Println("a % b  = ", a%b) // modulo de la variable a et b

	a++
	fmt.Println("a++ = ", a)
	a--

	fmt.Printf("a+c =", float32(a)+c)
}
