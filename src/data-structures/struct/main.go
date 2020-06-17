package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	yoda := sage{
		Name: "Yoda",
		Motto: "There is no try. Do or do not.",
	}

	musashi := sage{
		Name: "Miyamoto Musashi",
		Motto: "You can win with a long weapon, and yet you can also win with a short weapon. In short, " +
			"the Way of the Ichi school is the spirit of winning, whatever the weapon and whatever its size.",
	}

	jesus := sage{
		Name: "Jesus",
		Motto: "Do to others what you want them to do to you.",
	}

	// we need to create a var that holds a slice of our struct implementations; a slice of type sage
	sages := []sage{yoda, musashi, jesus}

	// here we're passing in the map named sages as the data
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
