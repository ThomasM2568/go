package main

import (
	"fmt"
)

func main() {
	var nombres = []int{0, 0, 0, 0, 0} // création d'une slice avec 5 éléments
	fmt.Println(nombres)

	var nombres2 = make([]int, 5) // création d'une slice avec 5 éléments
	fmt.Println(nombres2)

	nombres = append(nombres, 5)
	fmt.Println(nombres)

	nombres = append(nombres[:2], nombres[3:]...)
	fmt.Println(nombres)

	fmt.Println(nombres2)
	copy(nombres, nombres2)
	fmt.Println(nombres2)
}
