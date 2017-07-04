package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// this programm can be connected with terminal
// telnet localhost 8080
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

		fmt.Println("One user has connected")

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v \n", "WEEEEEEEEEEEEELLLL, I hope!!!!!")

		conn.Close()
	}
}
