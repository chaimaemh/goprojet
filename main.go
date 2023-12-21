package main

import (
	"fmt"
	"net/http"
	"goprojet/dictionnaire/handlers" 
)

func main() {
	// Créez un routeur pour gérer les différentes routes.
	mux := http.NewServeMux()

	// Associez les gestionnaires aux routes.
	mux.HandleFunc("/add", handlers.AddHandler)
	mux.HandleFunc("/get/", handlers.GetHandler)
	mux.HandleFunc("/remove/", handlers.RemoveHandler)
	mux.HandleFunc("/list", handlers.ListHandler)

	// Démarrez le serveur HTTP sur le port 8080.
	port := 8080
	fmt.Printf("Serveur démarré sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
