package main

//import obligatoire
//permet de dire que le code s'exécute directement dans le cas dde main

import "fmt"

//Permet de formater le texte

func main() {
	fmt.Println("Hello world !") //Afficher du texte
	/*
		Commentaire plus long
		sur plusieurs lignes
	*/
	var aaaa string = "aaaaaaaaaa"
	var (
		toto, tata, tutu int     = 10, 20, 30
		nom              string  = "hello"
		vitesse          float32 = 5.4
	)

	fmt.Println("Toto : ", toto)
	fmt.Println("Tata : ", tata)
	fmt.Println("Tutu : ", tutu)
	fmt.Println("nom : ", nom)
	fmt.Println("vitesse : ", vitesse)
	fmt.Println("aaaa : ", aaaa)

}
