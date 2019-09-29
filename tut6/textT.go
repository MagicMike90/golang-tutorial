package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need the template file!")
		return
	}

	tFile := arguments[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var Entries []Entry

	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Entries = append(Entries, temp)
		}
	}

	// return data type is Template, which is a structure that holds the representation of a parsed template
	t := template.Must(template.ParseGlob(tFile))
	// processing the data and printing the output to the desired file,
	t.Execute(os.Stdout, Entries)
}
