package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// run with go run main.go HeyFedja
func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
        <!DOCTYPE html>
        <html lang="de">
        <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
        </head>
        <body>
        <h1>` + name + ` </h1>
        </body>
        </html>
        `)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating index.html")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
