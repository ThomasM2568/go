package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func division() {
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entrez un chiffre : ")
		scanner.Scan()
		nbr, _ := strconv.Atoi(scanner.Text())
		if nbr <= 0 {
			fmt.Println("[division par zéro impossible] Votre valeur doit être supérieur ou égal à 0")
		} else {
			fmt.Println("Résultat :", 1000/nbr)
			break
		}
	}

}

func main() {
	division()
	fmt.Println("Fin")
}
