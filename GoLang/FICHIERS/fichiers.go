package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	//io
	//io/ioutil
)

type fichier struct {
	donnéesrecup string
}

func main() {
	var fichier1 fichier
	fichier1.donnéesrecup = ""

	menu(fichier1)
}

func menu(fichier1 fichier) {
	var choix int
	fmt.Println("Choissisez une option 1: recuperer le texte d'un fichier 2: ecrire dans un fichier 3: supprimer un fichier 4: remplacer le texte d'un fichier (n'oubliez pas de mettre les fichier en question dans le meme dossie que le programme)")
	fmt.Scanln(&choix)

	switch choix {

	case 1:
		{
			récupérer(fichier1)
		}

	case 2:
		{
			écrire(fichier1)
		}

	case 3:
		{
			supprimer(fichier1)
		}

	case 4:
		{
			remplacer(fichier1)
		}

	default:
		{
			fmt.Println("Vous avez entré une valeur invalide")
			menu(fichier1)
		}
		break
	}

}

func récupérer(fichier1 fichier) {
	var nom string

	fmt.Println("Entrez le nom du fichier (n'oubliez pas d'inclure l'extension du fichier)")
	fmt.Scanln(&nom)

	data, err := ioutil.ReadFile(nom) //lecture du contenu
	if err != nil {
		fmt.Println("Erreur : fichier introuvable")
		os.Exit(15)
	}

	bytesRead, err := ioutil.ReadFile(nom) //lecture du contenu

	if err != nil {
		panic(err)
	}
	var clip string //stockage du contenu dans une valeur
	clip = string(bytesRead)

	fmt.Println("valeurs récupérées")
	fmt.Scanln(data)

	fmt.Println(clip)
	fichier1.donnéesrecup = clip
	menu(fichier1)
}

func écrire(fichier1 fichier) string {
	var nomfichier string

	fmt.Println("entrez le nom du fichier (n'oubliez pas d'inclure l'extension du fichier)")
	fmt.Scanln(&nomfichier)

	str := "test"

	data, err := os.OpenFile(nomfichier, os.O_APPEND, 0644)
	if err != nil {
		os.Exit(15)
	}
	fmt.Println("entrez le texte à écrire (n'oubliez pas de d'appuyer sur entrée  puis controle + c pour enregister votre texte )")
	r, err1 := ioutil.ReadAll(os.Stdin)
	fmt.Scan(r)
	if err1 != nil {
		fmt.Println("erreur lors de l'écriture du fichier", err)
		os.Exit(2)
	}

	io.Copy(data, strings.NewReader(string(r)))
	//fmt.Println("'\n'Texte écrit")

	return str
}

func supprimer(fichier1 fichier) {
	var nomfichier string
	println("entrez le nom du fichier donc le contenu doit supprimé (n'oubliez pas d'inclure l'extension du fichier)")
	fmt.Scanln(&nomfichier)

	err := os.Remove(nomfichier)
	if err != nil {
		fmt.Println("erreur lors de la manipulation du fichier : ", err)
		os.Exit(1)
	}

	file, err := os.Create(nomfichier)
	if err != nil {
		fmt.Println("erreur lors de la manipulation du fichier : ", err)
		os.Exit(1)
	}
	defer file.Close()
	println("Contenu supprimé")
}

func remplacer(fichier1 fichier) {

	var nomfichier string

	println("entrez le nom du fichier donc le contenu doit etre remplacé (n'oubliez pas d'inclure l'extension du fichier)")
	fmt.Scanln(&nomfichier)

	err := os.Remove(nomfichier)
	if err != nil {
		fmt.Println("erreur lors de la manipulation du fichier : ", err)
		os.Exit(1)
	}

	file, err := os.Create(nomfichier)
	if err != nil {
		fmt.Println("erreur lors de la manipulation du fichier : ", err)
		os.Exit(1)
	}

	file, err = os.OpenFile(nomfichier, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("erreur de la manipulation du fichier : ", err)
		os.Exit(1)
	}

	println("entrez le texte à écrire (n'oubliez pas de d'appuyer sur entrée  puis controle + c pour enregister votre texte )")

	r1, err1 := ioutil.ReadAll(os.Stdin)
	fmt.Scan(r1)
	if err1 != nil {
		fmt.Println("erreur lors de l'écriture du fichier", err)
		os.Exit(2)
	}

	_, err = file.WriteString(string(r1))
	if err != nil {
		fmt.Println("erreur de l'ecriture du fichier : ", err)
		os.Exit(1)
	}

	defer file.Close()

}
