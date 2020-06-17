package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// ParseFiles method of text/template takes in zero or more string arguments, so we can pass in as many files as we want
	// we're creating a pointer to a template here with tpl, think of the pointer to the template as a container
	// in which all the templates we have parsed are held
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// create a new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()
	// Execute method of text/template takes a writer, we're using os.Stdout (Standard output) here, and passing in no data
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}