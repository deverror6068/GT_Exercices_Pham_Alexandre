package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email  string
	Nom    string
	Prenom string
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html")) // cible la page html hote

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:  r.FormValue("email"),
			Nom:    r.FormValue("nom"),
			Prenom: r.FormValue("prenom"),
		}

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})                 //la valeur Sucess determine si le formulaire est complété ou non
		fmt.Println("1re formulaire (email) : ", r.Form.Get("email")) // on recupère les information recupérées des formulaire et on les affiches sur la console
		fmt.Println("2ème formulaire (nom) : ", r.Form.Get("nom"))
		fmt.Println("3ème formulaire (prenom) : ", r.Form.Get("prenom"))
	})

	http.ListenAndServe(":80", nil) //on lance le serveur sur le port 80

}
