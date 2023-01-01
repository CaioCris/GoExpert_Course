package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	tmps := []string{
		"04_Pacotes_importantes/09_Templates/template_funcoes/header.html",
		"04_Pacotes_importantes/09_Templates/template_funcoes/content.html",
		"04_Pacotes_importantes/09_Templates/template_funcoes/footer.html",
	}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	tmp := template.Must(t.ParseFiles(tmps...))
	err := tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
