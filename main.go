package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request) {
	cardNum := r.FormValue("cardNumber")

	fmt.Println(cardNum)

	// read the number from string
	digits := []int{}
	counter := 0
	for _, digitCh := range cardNum {
		if digitCh >= '0' && digitCh <= '9' {
			digit, err := strconv.Atoi(string(digitCh))
			fmt.Println(digit)
			if err != nil {
				log.Fatal("Error: ", err)
			}
			digits = append(digits, digit)
			counter++
		} else {
			if counter > 16 {
				log.Fatal("Error: the count of numbers is more than 16 (unavailable format)")
			}
		}
	}

	var res bool

	if len(digits) != 16 {
		res = false
		tpl.ExecuteTemplate(w, "result.html", res)
		return
	}

	fmt.Println(digits)

	// luhn algo
	sum := 0
	reverse := false
	for i := len(digits) - 1; i >= 0; i-- {
		if reverse {
			digits[i] *= 2
			if digits[i] > 9 {
				digits[i] -= 9
			}
		}
		sum += digits[i]
		reverse = !reverse
	}

	res = sum%10 == 0

	fmt.Println(res)

	tpl.ExecuteTemplate(w, "result.html", res)
}

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/processGetHandler", processGetHandler)
	http.ListenAndServe(":8080", nil)
}
