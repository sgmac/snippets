package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/sys/unix"
)

var port int

func init() {
	flag.IntVar(&port, "p", 9090, "Port")
	flag.Parse()
}

func main() {
	buf := make([]byte, 50)
	fmt.Println("Listening on port :", port)

	// create socket
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("socket-error: ", err)
	}

	// sa struct
	sa := new(unix.SockaddrInet4)
	sa.Port = 9090

	// bind
	err = unix.Bind(fd, sa)
	if err != nil {
		log.Fatal("bind: ", err)
	}

	// listen
	err = unix.Listen(fd, 2)
	if err != nil {
		log.Fatal("listen: ", err)
	}

	for {
		// accept connection, discard SA struct
		newFd, _, err := unix.Accept(fd)
		if err != nil {
			log.Fatal("accept: ", err)
		}

		// client reads until closes
		for {
			// read
			n, err := unix.Read(newFd, buf)
			if err != nil {
				log.Fatal("read: ", err)
			}

			fmt.Printf("Read: %d Value: %s\n", n, string(buf[0:n]))

			// close
			if string(buf[0:5]) == "close" {
				_, err = unix.Write(newFd, []byte(`Bye bye buddy`))
				if err != nil {
					log.Fatal("close: ", err)
				}

				err = unix.Close(newFd)
				if err != nil {
					log.Fatal("close: ", err)
				}
			}
		}
	}
}
