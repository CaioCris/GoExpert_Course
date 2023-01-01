package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	tmps := []string{
		"04_Pacotes_importantes/09_Templates/template_compose/header.html",
		"04_Pacotes_importantes/09_Templates/template_compose/content.html",
		"04_Pacotes_importantes/09_Templates/template_compose/footer.html",
	}
	t, err := template.New("content.html").ParseFiles(tmps...)
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
