package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type PageVariables struct {
	Title string
	Word  string
	Sum   int
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/calculate", Calculate)
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func Index(w http.ResponseWriter, r *http.Request) {

	IndexVars := PageVariables{
		Title: "Let's Add Up Word Values!",
	}

	t, err := template.ParseFiles("index.html") // parse the html file
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// execute the template and pass it the IndexVars struct
	err = t.Execute(w, IndexVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	word := getWord(w, r)
	sum := convert(word)

	CalculateVars := PageVariables{
		Title: "Word Calculation Complete",
		Word:  word,
		Sum:   sum,
	}

	t, err := template.ParseFiles("calculate.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, CalculateVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func getWord(w http.ResponseWriter, r *http.Request) string {
	var word string
	r.ParseForm()
	for _, v := range r.Form {
		word = strings.Join(v, "")
	}
	return word
}

func convert(s string) int {
	s = strings.ToLower(s)
	var i, j int
	for _, c := range s {
		switch string(c) {
		case "a":
			j = 1
		case "b":
			j = 2
		case "c":
			j = 3
		case "d":
			j = 4
		case "e":
			j = 5
		case "f":
			j = 6
		case "g":
			j = 7
		case "h":
			j = 8
		case "i":
			j = 9
		case "j":
			j = 10
		case "k":
			j = 11
		case "l":
			j = 12
		case "m":
			j = 13
		case "n":
			j = 14
		case "o":
			j = 15
		case "p":
			j = 16
		case "q":
			j = 17
		case "r":
			j = 18
		case "s":
			j = 19
		case "t":
			j = 20
		case "u":
			j = 21
		case "v":
			j = 22
		case "w":
			j = 23
		case "x":
			j = 24
		case "y":
			j = 25
		case "z":
			j = 26
		default:
			j = 0
		}
		i = i + j
	}
	return i
}
