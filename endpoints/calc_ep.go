package endpoints

import (
	"html/template"
	"net/http"
	"strconv"
	calc "cicd-pipeline-go/utils/calculator"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
				res = calc.Sum(num1, num2)
			case "restar":
				res = calc.Sub(num1, num2)
			case "multiplicar":
				res = calc.Mult(num1, num2)
			case "dividir":
				ress, err := calc.Div(num1, num2)
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