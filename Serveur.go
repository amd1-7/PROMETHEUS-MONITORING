package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	log.Printf("Initialisation du serveur...\n")

	if err := http.ListenAndServe(":1221", mux); err != nil {
		log.Fatalf("Erreur de lancement: %v\n", err)
	}
}
