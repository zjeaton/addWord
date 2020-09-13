package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Title string
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
