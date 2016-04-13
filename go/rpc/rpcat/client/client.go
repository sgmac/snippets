package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/sgmac/rpcat"
)

func main() {

	var response rpcat.Cmd
	fileName := flag.String("f", "", "filename to remote cat")
	flag.Parse()

	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("RPC dialing..")
	client, err := rpc.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Println("RPC call...")
	err = client.Call("Cmd.Cat", *fileName, &response)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println(response.Output)

}
