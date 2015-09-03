package main

import (
	"bufio"
	"log"
	"net"
)

var banner string = `
 _____     _
 | ____|___| |__   ___
 |  _| / __| '_ \ / _ \
 | |__| (__| | | | (_) |
 |_____\___|_| |_|\___/)'
		 ____
		 / ___|  ___ _ ____   _____ _ __
		 \___ \ / _ \ '__\ \ / / _ \ '__|
		  ___) |  __/ |   \ V /  __/ |
		  |____/ \___|_|    \_/ \___|_|
`

func main() {
	log.Println("Listening on :9090")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		conn.Write([]byte("Welcome to \n"))
		conn.Write([]byte(banner))
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			for {
				reader := bufio.NewReader(c)
				line, _ := reader.ReadString(byte('\n'))
				c.Write([]byte(string(line)))
			}
		}(conn)
	}
}
