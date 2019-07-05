package main

import (
	"html/template"
	"net/http"
)

func process(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/tmpl.html")
	tmpl.Execute(writer, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", process)
	server.ListenAndServe()
}
