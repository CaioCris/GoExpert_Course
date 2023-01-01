package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.New("template.html").ParseFiles("04_Pacotes_importantes/09_Templates/template_webserver/template.html")
		tmp := template.Must(t, err)
		err = tmp.Execute(writer, Courses{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8082", nil)

}
