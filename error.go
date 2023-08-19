package main

import (
	"errors"
	"fmt"
)

func verificationDivision(nbr float64) (float64, error) {
	if nbr <= 0 {
		return 0, errors.New("Erreur: Il est impossible de diviser par 0 !") //création de l'erreur
	} else {
		return nbr, nil // on retourne nil si aucune erreur est détectée
	}
}

func main() {
	nbr := 0.0

	nbr, err := verificationDivision(nbr)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Aucune erreur")
	}
}
