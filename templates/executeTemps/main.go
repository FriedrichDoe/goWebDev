package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// parsing data into template
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", `GOLANG IS AWESOME!!!`)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(time.Now().Format("01.02.2006"))
}
