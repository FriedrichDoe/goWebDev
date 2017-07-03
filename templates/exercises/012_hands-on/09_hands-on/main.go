package main

import (
	"encoding/csv"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type table struct {
	Date string
	Open string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	tables := make([]table, 0)

	for i := 1; i < len(records); i++ {
		n := table{
			records[i][0],
			records[i][1],
		}
		tables = append(tables, n)
	}

	errr := tpl.Execute(res, tables)
	if errr != nil {
		log.Fatalln(err)
	}
}
