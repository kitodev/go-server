package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		
	}
}