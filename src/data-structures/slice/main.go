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
	sages := []string{"Yoda", "Ghandi", "MLK", "Buddah", "Miyamoto Musashi", "Jesus", "Mohammed"}

	// here we're passing in the slice named sages as the data
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
