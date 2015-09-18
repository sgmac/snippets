package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/Sirupsen/logrus"
)

func main() {
	log.Println("Listening on :7070")
	var a, b, result int
	var op string
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		conn.Write([]byte("Welcome to \n"))
		conn.Write([]byte("CalCulator\n"))
		conn.Write([]byte(">"))
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			for {
				reader := bufio.NewReader(c)
				line, _ := reader.ReadString(byte('\n'))
				_, err := fmt.Sscanf(line, "%d %s %d", &a, &op, &b)
				if err != nil {
					c.Write([]byte(err.Error() + "\n"))
					logrus.Fatal(err)
				}
				switch op {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					if b != 0 {
						result = a / b
					}
				}
				c.Write([]byte(strconv.Itoa(result) + "\n"))
				conn.Write([]byte(">"))
				// reset result otherwise shows the previous one
				result = 0
			}
		}(conn)
	}
}
