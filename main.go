package main

import (
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	data := Images{
		Image:images[rand.Intn(len(images))],
	}
	tpl.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// Add the following two lines
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}