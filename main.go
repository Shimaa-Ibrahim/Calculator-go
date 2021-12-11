package main

import (
	"html/template"
	"learn-go/calculator"
	"log"
	"net/http"
	"strconv"
)

// var num1, num2 float64
// var operator string

type htmlCalculator struct {
	Fnum, Snum, Result     float64
	Operator, ErrorMessage string
	IsValid                bool
}

var calc htmlCalculator

func handler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/calculator.html")
	err := template.Execute(w, calc)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err1, err2 error
	calc.Fnum, err1 = strconv.ParseFloat(r.FormValue("fnum"), 64)
	calc.Operator = r.FormValue("operator")
	calc.Snum, err2 = strconv.ParseFloat(r.FormValue("snum"), 64)

	isOperator := calculator.CheckOperator(calc.Operator)

	if err1 != nil || err2 != nil || !isOperator {
		calc.IsValid = false
		calc.ErrorMessage = "Ivalid Operation! please try again"
	} else {
		calc.IsValid = true
		calc.ErrorMessage = ""
		calc.Result = calculator.Calculate(calc.Fnum, calc.Operator, calc.Snum)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/calc", calcHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/* console prog
func main() {
	printAndScan("Enter the first number : ", &num1)
	printAndScan("Enter your operator : ", &operator)
	printAndScan("Enter the second number : ", &num2)
	calculator.Calculate(num1, operator, num2)
}
func printAndScan(msg string, variable interface{}) {
	fmt.Print(msg)
	_, err := fmt.Scan(variable)
	if err != nil {
		fmt.Print("Invalid number please try again... ", err)
		printAndScan(msg, variable)
	}
}
*/
