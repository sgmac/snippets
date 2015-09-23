package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

var certPath string

func init() {
	flag.StringVar(&certPath, "c", "", "Certificate file")
	flag.Parse()
}

func main() {

	if flag.NFlag() < 1 {
		fmt.Fprintf(os.Stderr, "usage: vecert [options] certfile\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	byteCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// decode pem
	p, _ := pem.Decode(byteCert)
	cert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("CN=%s\n", cert.Subject.CommonName)
	fmt.Printf("Country=%s\n", cert.Subject.Country)
	fmt.Printf("Organization=%s\n", cert.Subject.Organization)

}
