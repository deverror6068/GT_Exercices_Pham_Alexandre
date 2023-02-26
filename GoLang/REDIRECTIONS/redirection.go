package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var home = template.Must(template.ParseFiles("home.html"))
var page2 = template.Must(template.ParseFiles("page2.html"))
var page3 = template.Must(template.ParseFiles("page3.html"))

func redirection1(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)

	if r.URL.Path == "/" {
		http.Redirect(w, r, "/page2", http.StatusSeeOther)
		fmt.Println(r.URL.Path)
		return
	}

	if r.URL.Path == "/page2" {
		http.Redirect(w, r, "/page3", http.StatusSeeOther)
		fmt.Println(r.URL.Path)
		return
	}

	if r.URL.Path != "/page3" {

		fmt.Println(r.URL.Path)
		errorHandler(w, r, http.StatusNotFound)
		return
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	fmt.Fprintln(w, " Error 404 page not found")
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	data := 50
	home.Execute(w, data)
}

func HttpHandler2(w http.ResponseWriter, r *http.Request) {
	data := 50
	page2.Execute(w, data)
}

func HttpHandler3(w http.ResponseWriter, r *http.Request) {
	data := 50
	page3.Execute(w, data)
}

func main() {

	mux := http.NewServeMux()
	home = template.Must(template.ParseFiles("home.html"))
	page2 = template.Must(template.ParseFiles("page2.html"))
	page3 = template.Must(template.ParseFiles("page3.html"))

	mux.HandleFunc("/", redirection1)

	http.HandleFunc("/", HttpHandler)
	http.HandleFunc("/page2", HttpHandler2)
	http.HandleFunc("/page3", HttpHandler3)

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", mux)
}
