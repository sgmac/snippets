package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/sgmac/rpcat"
)

func main() {

	command := new(rpcat.Cmd)
	rpc.Register(command)
	l, err := net.Listen("tcp", "127.0.0.1:3000")
	log.Println("Serving connetions on :3000")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}
		rpc.ServeConn(conn)
		conn.Close()
	}

}
