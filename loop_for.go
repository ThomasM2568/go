package main

import (
	"fmt"
)

func main() {
	for compteur := 0; compteur < 100; compteur++ {
		fmt.Println(compteur+1, ") Je ne dois frapper mes camarades de classe")
		switch compteur {
		case 10:
			println("Et de 10 !")
		case 99:
			println("Allez encore une et on est bon !")
		}

		switch {
		case compteur == 50:
			println("On est à la moitié !")
		case compteur == 80:
			println("Pas loin de la fin !")
		}

		if compteur == 60 {
			println("60 !")
		} else if compteur == 20 {
			println("C'est bien on avance !")
		}
	}
}
