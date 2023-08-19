package main

import "fmt"

type Player struct {
	nom        string
	vie        int
	puissance  int
	mort       bool
	inventaire [3]string
}

func valeurParDefaut(p *Player) {
	p.nom = "inconnu"
	p.vie = 50
	p.puissance = 10
	p.mort = false
	p.inventaire = [3]string{"vide", "vide", "vide"}
}

func (p Player) affichage() {
	fmt.Println("Le personnage s'appelle : ", p.nom)
	fmt.Println("Nombre de points de vie : ", p.vie)
	fmt.Println("Nombre de points de puissance : ", p.puissance)
	if !p.mort {
		fmt.Println("Le personnage n'est pas mort")
	} else {
		fmt.Println("Le personnage est mort")
	}
	fmt.Println("L'inventaire contient :")
	for _, item := range p.inventaire {
		fmt.Println("- ", item)
	}
}

func (p *Player) Init(nom string, vie int, puissance int, mort bool, inventaire [3]string) {
	p.nom = nom
	p.vie = vie
	p.puissance = puissance
	p.mort = mort
	p.inventaire = inventaire
}

func main() {

	// déclaration de la structure nommée Personnage
	type Personnage struct {
		nom string
		age int
	}

	perso := Personnage{"Hatim", 20} // surchargement des valeurs par défaut
	fmt.Println(perso)
	perso.age = 32
	fmt.Println(perso)

	var test Player
	test.nom = "aa"
	test.vie = 45
	test.puissance = 123
	test.mort = false
	test.inventaire[0] = "Ordinateur"

	test.affichage()

}
