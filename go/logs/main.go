package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var InfoLogger *log.Logger

func main() {
	fmt.Println("Managing logs")

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logfile := io.MultiWriter(file, os.Stdout)

	InfoLogger = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("command not found")
}
