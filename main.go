package main

import (
	"fmt"
	"html/template"
	"learn/shape"
	"net/http"
	"time"
)

type Welcome struct {
	Detail string
	Name   string
	Time   string
}

func main() {
	welcome := Welcome{"", "user", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("template/welcome-template.html"))
	// var r shape.Shape = shape.Rect{10, 3}
	var c shape.Shape = shape.Circle{5}

	welcome.Detail = shape.Detail(c)
	http.Handle("/statics/",
		http.StripPrefix("/statics/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
