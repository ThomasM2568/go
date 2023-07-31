/*
package main

import (

	"bufio"
	"fmt"
	"os"
	"strconv"

)

	func main() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Votre choix : ")
		scanner.Scan()
		choix, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Entrez un entier !")
			os.Exit(2)
		}

		switch choix { // la variable qu'on souhaite vérifier
		case 0, 1: // 1 ou 0
			println("George Boole !")
		case 7:
			println("William Van de Walle!")
		case 23:
			println("Pour certains, le nombre 23 est source de nombreux mystères, qu'en penses-tu Jim Carrey?")
		case 42:
			println("la réponse à la question ultime du sens de la vie!")
		case 666:
			println("Quand le diable veut une âme, le mal devient séduisant.")
		default:
			println("Mauvais choix !") // Valeur par défaut
		}
	}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrez un entier !")
		os.Exit(2)
	}

	switch {
	case choix >= 2000:
		println("Ah un 2000 !")
	case choix >= 1939 && choix <= 1945:
		println("Triste année")
	case time.Now().Weekday() == time.Sunday:
		println("Dimanche !")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}
}
