package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"golang.org/x/sys/unix"
)

var port int

type connection struct {
	buf []byte
	fd  int
}

type connCounter struct {
	num int
	mu  sync.Mutex
}

func init() {
	flag.IntVar(&port, "p", 9090, "Port")
	flag.Parse()
}

func main() {
	fmt.Println("Listening on port :", port)
	counter := new(connCounter)

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
	err = unix.Listen(fd, 3)
	if err != nil {
		log.Fatal("listen: ", err)
	}

	for {
		connInfo := new(connection)
		// accept connection, discard SA struct
		newFd, _, err := unix.Accept(fd)
		connInfo.fd = newFd
		if err != nil {
			log.Fatal("accept: ", err)
		}

		// client reads until closes, adding
		// a gorutine allows dealing with more
		// requests.
		counter.mu.Lock()
		counter.num += 1
		counter.mu.Unlock()
		fmt.Println("Number of connections: ", counter.num)
		go func(c *connection, counter *connCounter) {
			fmt.Printf("Conn.fd=%d\n", c.fd)
			for {
				// read
				c.buf = make([]byte, 50)
				n, err := unix.Read(c.fd, c.buf)
				if err != nil {
					log.Fatal("read: ", err)
				}

				fmt.Printf("Read: %d Value: %s\n", n, string(c.buf[0:n]))

				// close
				if string(c.buf[0:5]) == "close" {
					_, err = unix.Write(c.fd, []byte(`Bye bye buddy`))
					if err != nil {
						log.Fatal("close: ", err)
					}

					err = unix.Close(c.fd)
					if err != nil {
						log.Fatal("close: ", err)
					}
					counter.mu.Lock()
					counter.num = counter.num - 1
					counter.mu.Unlock()
					return
				}
			}
		}(connInfo, counter)
	}
}
