package main

import (
	"datecalculate/src/apidatecalculate"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.HandleFunc("/duration/", apidatecalculate.ApiDateCalculate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
