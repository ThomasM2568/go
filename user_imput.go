/*
package main

import (

	"bufio"
	"fmt"
	"os"
	"strconv"

)

	func main() {
		/*scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
		fmt.Print("Entrez quelque chose : ")
		scanner.Scan()                      // lancement du scanner
		entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
		fmt.Println(entreeUtilisateur)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entrez un nombre entier : ")
		scanner.Scan()

		nbr, _ := strconv.Atoi(scanner.Text()) // conversion du type string en int

		fmt.Printf("res : %d\n", (nbr + 6))

}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
		fmt.Println("On essaie de m'arnaquer ? allé Dehors ! Et la prochaine entrez un entier !")
		os.Exit(2) // on quitte le programmation
	}

	if age < 17 && age > 90 { // vérifier si l'utilisateur à au moins 18 ans et plus de 90
		fmt.Println("if !")
	} else if age == 18 { // si l'age est 18
		fmt.Println("else if :)")
	} else {
		fmt.Println("else")
	}
}
