package main

import (
	"log"
	"net/http"
	"cicd-pipeline-go/endpoints"
)

func main() {
	http.HandleFunc("/", endpoints.IndexHandler)
	log.Println("Servidor escuchando en http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
