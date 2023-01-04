package main

import (
	"html/template"
	"os"
)

type course struct {
	Name     string
	Workload int
}

func main() {
	course := course{"Go", 40}
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}} - Carga Horária: {{.Workload}}")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
