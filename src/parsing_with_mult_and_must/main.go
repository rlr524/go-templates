package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// when we ParseGlob we get back a pointer to a template (tpl) and an err
// Must then takes in a pointer to a template and an error, and it handles errors returned by our tpl pointer
// to template.Template; we do not need explicit error handling in our init function
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//// keeping this also provides an extra output of angela.gohtml as it processes the templates from the func init
	//// above in alphabetical order; we can remove it and as long as our first execute uses the type coercion operator
	//// it will be ok
	//err := tpl.Execute(os.Stdout, nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	err := tpl.ExecuteTemplate(os.Stdout, "angela.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	_ = tpl.ExecuteTemplate(os.Stdout, "madison.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	_ = tpl.ExecuteTemplate(os.Stdout, "rila.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
