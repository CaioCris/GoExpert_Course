package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	t, err := template.New("template.html").ParseFiles("04_Pacotes_importantes/09_Templates/template_file/template.html")
	tmp := template.Must(t, err)
	err = tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
