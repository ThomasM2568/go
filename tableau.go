package main

import (
	"fmt"
)

func main() {
	var tableauInt [10]int
	var tableauFloat [10]float32
	var tableauString [10]string
	var tableauBool [10]bool
	var tableau1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("la taille de mon tableau1 :", len(tableau1))
	fmt.Println("les valeurs de mon tableau1 :", tableau1)
	fmt.Println("Valeur par défaut de la variable tableauInt :", tableauInt)
	fmt.Println("Valeur par défaut de la variable tableauFloat :", tableauFloat)
	fmt.Println("Valeur par défaut de la variable tableauString :", tableauString)
	fmt.Println("Valeur par défaut de la variable tableauBool :", tableauBool)

	fmt.Println("Valeur 1 du tableau : ", tableau1[1])
	fmt.Println("Dernière valeur du tableau : ", tableau1[len(tableau1)-1])
	var a int = 10
	tableau1[1] = a
	fmt.Println("Valeur 1 du tableau : ", tableau1[1])

	for j := range tableau1 {
		fmt.Println(tableau1[j])
	}

	for i := 0; i < len(tableau1); i++ {
		fmt.Println(tableauBool[i])
	}

	fmt.Println(tableau1[:2])
	fmt.Println(tableau1[2:])
	fmt.Println(tableau1[1:3])
	var tableau [2][3]int // Création d'un tableau à double dimension

	fmt.Println(tableau)
}
