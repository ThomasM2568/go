package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func game(tableau [9]string, pos int, player string) [9]string {
	if !win(tableau) {
		if tableau[pos] == "X" || tableau[pos] == "O" {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("la case est déjà occupée, veuillez entrez un autre entier !")
			scanner.Scan()
			npos, err := strconv.Atoi(scanner.Text())
			if err != nil {
				// Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
				fmt.Println("On essaie de m'arnaquer ? allé Dehors ! Et la prochaine entrez un entier !")
			}
			return game(tableau, npos, player)
		} else {
			tableau[pos] = player
			return tableau
		}
	} else {
		var tableau2 [9]string
		return tableau2
	}
}

func new_game() [9]string {
	var tableau [9]string
	for j := 0; j < len(tableau); j++ {
		tableau[j] = strconv.Itoa(j)
	}

	return tableau
}

func win(tableau [9]string) bool {
	// Vérifier les lignes
	for i := 0; i <= 6; i += 3 {
		if tableau[i] != "" && tableau[i] == tableau[i+1] && tableau[i+1] == tableau[i+2] {
			return true
		}
	}

	// Vérifier les colonnes
	for i := 0; i < 3; i++ {
		if tableau[i] != "" && tableau[i] == tableau[i+3] && tableau[i+3] == tableau[i+6] {
			return true
		}
	}

	// Vérifier les diagonales
	if tableau[0] != "" && tableau[0] == tableau[4] && tableau[4] == tableau[8] {
		return true
	}
	if tableau[2] != "" && tableau[2] == tableau[4] && tableau[4] == tableau[6] {
		return true
	}

	return false
}

func switch_player(player string) string {
	switch player {
	case "X":
		return "O"
	case "O":
		return "X"
	}
	return "toto"
}

func main() {
	var tictactoe [9]string = new_game()

	player := "X" // Le joueur X commence

	for {
		fmt.Println("Avancement du plateau :")
		for i := 0; i < len(tictactoe); i += 3 {
			fmt.Println(tictactoe[i], "|", tictactoe[i+1], "|", tictactoe[i+2])
			if i < len(tictactoe)-3 {
				fmt.Println("---------")
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Joueur", player, ", entrez un entier de 0 à 8 pour jouer :")
		scanner.Scan()
		pos, err := strconv.Atoi(scanner.Text())
		if err != nil || pos < 0 || pos > 8 {
			// Si la conversion a échoué ou que la position n'est pas valide, affichez un message d'erreur
			fmt.Println("Position invalide. Entrez un entier entre 0 et 8.")
			continue
		}

		// Appel de la fonction game pour mettre à jour le plateau avec le coup du joueur
		tictactoe = game(tictactoe, pos, player)

		// Vérifier si le joueur actuel a gagné
		if win(tictactoe) {
			fmt.Println("Le joueur", player, "a gagné !")
			break
		}

		// Vérifier si le plateau est rempli (match nul)
		boardFilled := true
		for _, cell := range tictactoe {
			if cell != "X" && cell != "O" {
				boardFilled = false
				break
			}
		}
		if boardFilled {
			fmt.Println("Match nul !")
			break
		}

		// Passer au joueur suivant
		player = switch_player(player)
	}
}
