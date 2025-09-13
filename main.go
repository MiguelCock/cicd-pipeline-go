package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var res interface{} = nil

	if r.Method == http.MethodPost {
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")
		operacion := r.FormValue("operacion")

		num1, err1 := strconv.ParseFloat(num1Str, 64)
		num2, err2 := strconv.ParseFloat(num2Str, 64)

		if err1 != nil || err2 != nil {
			res = "error: enter valid numbers"
		} else {
			switch operacion {
			case "sumar":
				res = sum(num1, num2)
			case "restar":
				res = sub(num1, num2)
			case "multiplicar":
				res = mult(num1, num2)
			case "dividir":
				ress, err := div(num1, num2)
				if err != nil {
					res = err.Error()
				} else {
					res = ress
				}
			default:
				res = "error: invalid operation"
			}
		}
	}

	err := tmpl.Execute(w, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Servidor escuchando en http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
