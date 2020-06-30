package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

type Anime struct {
	Title       string
	Description string
}

type Character struct {
	Name    string
	Species string
	Age     int
}

// create a FuncMap to register functions
// "uc" us what the func will be called in the template
// "uc" is the topper func from packaged strings
// "ft" is a func we declared
// "ft" slices a string, returning the first three chars

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	//"fdateMDY": monthDayYear,
}

func init() {
	// we use chaining here because we need to run our funcs before we parse the files
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
//// see the constants in the time package in the std lib to see the constants that can be used aside from RFC1123
//func monthDayYear(t time.Time) string {
//	return t.Format(time.RFC1123)
//}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	k := Anime{
		Title:       "K-On!",
		Description: "Five girls play in a high school rock band and drink tea",
	}

	a := Anime{
		Title:       "Attack on Titan",
		Description: "Giant people, some of who look like Nixon, attack much smaller people and the Survey Corps must fight them",
	}

	h := Anime{
		Title:       "The Melancholy of Haruhi Suzumiya",
		Description: "God is a pushy 15-year-old Japanese girl and you better not bore her",
	}

	y := Character{
		Name:    "Yui Hirasawa",
		Species: "Human, really, really moe human",
		Age:     16,
	}

	m := Character{
		Name:    "Mikasa Ackerman",
		Species: "Human, but far more badass than you",
		Age:     16,
	}

	n := Character{
		Name:    "Yuki Nagato",
		Species: "Alien android supercomputer",
		Age:     15,
	}

	shows := []Anime{k, a, h}
	girls := []Character{y, m, n}

	var data = struct {
		Works []Anime
		Stars []Character
	}{
		shows,
		girls,
	}

	//err := tpl.ExecuteTemplate(os.Stdout, "tplTime.gohtml", time.Now())
	//if err != nil {
	//	log.Fatalln(err)
	//}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
