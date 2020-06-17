package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"India": "Ghandi",
		"America": "MLK",
		"Meditate": "Buddah",
		"Love": "Jesus",
		"Prophet": "Muhammad",
		"The Way": "Miyamoto Musashi",
		"There is no try": "Yoda",
	}

	// here we're passing in the map named sages as the data
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}