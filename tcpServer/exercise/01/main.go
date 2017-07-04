package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// gets the url from the browser
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		url := strings.Fields(ln)[1]
		fmt.Println("URL called:", url)
		break
	}
}
