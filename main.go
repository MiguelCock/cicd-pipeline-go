package main

import (
	"cicd-pipeline-go/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoints.IndexHandler)
	log.Println("Servidor escuchando en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
