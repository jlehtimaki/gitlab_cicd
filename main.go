package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var tpl = template.Must(template.ParseFiles("index.html"))
var images = []string {
	"jeremy.png",
}

type Images struct {
	Image	string
}

type Status struct {
	Status		int		`json:"status"`
	Response 	string	`json:"response"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	data := Images{
		Image:images[rand.Intn(len(images))],
	}
	tpl.Execute(w, data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	returnData := Status{
		Status:   200,
		Response: "OK",
	}
	json.NewEncoder(w).Encode(returnData)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Add the following two lines
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":"+port, mux)
}