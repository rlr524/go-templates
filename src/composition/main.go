package main

import (
	"html/template"
	"log"
	"os"
)

type Course struct {
	Number, Name, Units string
}

type Semester struct {
	Term string
	Courses []Course
}

type Year struct {
	Fall, Spring, Summer Semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := Year {
		Fall: Semester{
			Term:    "Fall",
			Courses: []Course{
				{
					Number: "CSCI-40",
					Name:   "Introduction to Programming in Go",
					Units:  "4",
				},
				{
					Number: "CSCI-130",
					Name:   "Introduction to Web Programming with Go",
					Units:  "4",
				},
				{
					Number: "CSCI-140",
					Name:   "Mobile Apps Using Go",
					Units:  "4",
				},
			},
		},
		Spring: Semester{
			Term:    "Spring",
			Courses: []Course {
				{
					Number: "CSCI-50",
					Name:   "Advanced Go",
					Units:  "5",
				},
				{
					Number: "CSCI-190",
					Name:   "Advanced Web Programming with Go",
					Units:  "5",
				},
				{
					Number: "CSCI-191",
					Name:   "Advanced Mobile Apps with Go",
					Units:  "5",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
