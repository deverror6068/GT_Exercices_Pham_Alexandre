package main

import (
	"fmt"
	"os"
	"reflect"
)

type Jeu struct {
	nballumettes        int //nombre d'alumettes
	nballumettespartour int //nombre d'allumettes choisi
	allumettes          int //allumettes restantes
	tours               int
	nbjoueurs           int
	autourdujoueur      int //le tour du joueur
}

func main() {
	var Jeu1 Jeu
	var int int
	Jeu1.allumettes = 0
	Jeu1.nballumettes = 0
	Jeu1.nballumettespartour = 0
	Jeu1.tours = 0
	Jeu1.nbjoueurs = 2
	Jeu1.autourdujoueur = 0

	fmt.Println(" \nEntrez le nombre d'allumettes (au minimum 4)")
	fmt.Scanln(&Jeu1.nballumettes)

	if Jeu1.nballumettes <= 4 || reflect.TypeOf(Jeu1.nballumettes) != reflect.TypeOf(int) { // si l'utilisateur a entré une valeur qui n'est pas un nombre  ou une valeur inférieure à 4

		fmt.Println(" Vous devez entrer un entier et il faut au moins 4 allumettes \n \n ")

		for Jeu1.nballumettes <= 3 || reflect.TypeOf(Jeu1.nballumettes) != reflect.TypeOf(int) { // tant que  l'utilisateur  n'a pas  entré une valeur correcte
			fmt.Println("\nEntrez le nombre d'allumettes")
			fmt.Scanln(&Jeu1.nballumettes)
		}

	}

	Jeu1.allumettes = Jeu1.nballumettes //la variable prend la valeur entrée par le joueur

	tour(Jeu1) // appel de la fonction qui s'occupe des tours

}

func ver(Jeu1 Jeu) {

	if Jeu1.allumettes <= 1 { //condition de victoire / défaite
		Jeu1.tours += 1
		Jeu1.autourdujoueur += 1
		if Jeu1.autourdujoueur == Jeu1.nbjoueurs+1 {
			Jeu1.autourdujoueur = 1
		}
		print("il reste une allumette")
		print("\nLe joueur", Jeu1.autourdujoueur, " a perdu au tour  ", Jeu1.tours)

		os.Exit(0)
	}
	tour(Jeu1) // appel de la fonction qui s'occupe des tours

}

func tour(Jeu1 Jeu) {

	var int int
	Jeu1.tours += 1
	Jeu1.autourdujoueur += 1 // au tour du joueur suivant
	if Jeu1.autourdujoueur == Jeu1.nbjoueurs+1 {
		Jeu1.autourdujoueur = 1
	}
	fmt.Println("\nC'est au tour du joueur", Jeu1.autourdujoueur) //si l'utilisateur a entré une valeur qui n'est pas un nombre

	fmt.Println("\nEntrez le nb d'allumettes")
	fmt.Scanln(&Jeu1.nballumettespartour)
	if reflect.TypeOf(Jeu1.nballumettespartour) != reflect.TypeOf(int) {

		fmt.Println("\nVous devez entrez un entier\n ")

		for reflect.TypeOf(Jeu1.nballumettespartour) != reflect.TypeOf(int) {
			fmt.Println("\nEntrez le nombre d'allumettes1")
			fmt.Scanln(&Jeu1.nballumettes)
		}

	}

	if Jeu1.nballumettespartour < 1 || Jeu1.nballumettespartour > 3 { //si le joueur ne choisit pas un nombre choisi entre 1 et 3

		fmt.Println("\nVotre nombre doit etre compris entre 1 et 3")

		for Jeu1.nballumettespartour < 1 || Jeu1.nballumettespartour > 3 {

			fmt.Println("\nEntrez un nombre d'allumettes compris entre 1 et 3")
			fmt.Scanln(&Jeu1.nballumettespartour)
		}
	}

	if Jeu1.allumettes <= Jeu1.nballumettespartour {
		fmt.Println("Vous perdez si vous faites ca !") //empeche le joueur de perdre délibérémement
		for Jeu1.nballumettespartour >= Jeu1.allumettes {
			fmt.Println("\nEntrez un nombre d'allumettes qui n'est pas supérieur au nombre d'allumettes restantes")
			fmt.Scanln(&Jeu1.nballumettespartour)

		}

	}
	Jeu1.allumettes -= Jeu1.nballumettespartour

	ver(Jeu1) //appel de la fonction de verification de victoire
}
