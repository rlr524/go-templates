package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// each source file can define its own niladic init function (niladic meaning it returns nothing)
// to set up whatever state is required. init is a built-in function and does not need to be explicitly called
// (Actually each file can have multiple init functions.) And finally means finally: init is called
// after all the variable declarations in the package have evaluated their initializers, and those are
// evaluated only after all the imported packages have been initialized.
// besides initializations that cannot be expressed as declarations, a common use of init functions
// is to verify or repair correctness of the program state before real execution begins.
func init() {
	// using ParseFiles here vs ParseGlob because we're only parsing one file, not a folder of files
	tpl = template.Must(template.ParseFiles("tplnew.gohtml"))
}
// we can only pass in one piece of data but it can be of an aggregate type / data structure, such as a struct
// or a slice or a map in our below examples we're only passing in single values of types int or string
func main() {
	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	// call our data here from the var in tplnew
	err := tpl.ExecuteTemplate(os.Stdout, "tplnew.gohtml", "Try not. Do. Or do not. There is no try.")
	if err != nil {
		log.Fatalln(err)
	}
}
