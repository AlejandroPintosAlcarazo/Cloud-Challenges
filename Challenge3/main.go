package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ingles", inglesHandler)
	http.HandleFunc("/frances", francesHandler)
	svr := http.Server{
		Addr: ":8080",
	}
	
	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func inglesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there visitor/n")
}
func francesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour les visiteurs/n") // Frase en francés
}
