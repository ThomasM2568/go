package main

import "fmt"

func main() {
	flt := 15.5
	in := 5
	st := "hello"
	bol := true
	fmt.Printf("Le type de la variable flt est %T\n", flt)
	fmt.Printf("Le type de la variable in est %T\n", in)
	fmt.Printf("Le type de la variable st est %T\n", st)
	fmt.Printf("Le type de la variable bol est %T\n", bol)

	const maConstante int = 50
	fmt.Println("Ma constante : ", maConstante)
	/*

	   %T : affiche le type d'une valeur
	   %d : affiche un entier
	   %s : affiche une chaîne de caractères
	   %f : affiche un nombre décimal
	   %b : affiche une représentation binaire

	*/
}
