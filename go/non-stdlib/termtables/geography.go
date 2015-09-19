package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/apcera/termtables"
)

const ApiURL = "https://restcountries.eu/rest/v1/all"

type Country struct {
	Name       string
	Capital    string
	Region     string
	Population int
}

type Response []Country

func listContries() (Response, error) {
	var c Response
	r, err := http.Get(ApiURL)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func main() {
	fmt.Println("Listing countries...")
	// TODO: sort by any column
	countries, err := listContries()
	if err != nil {
		logrus.Fatal(err)
	}
	table := termtables.CreateTable()
	table.AddHeaders("Country", "Capital", "Region", "Population")

	for _, c := range countries {
		table.AddRow(c.Name, c.Capital, c.Region, c.Population)
	}
	fmt.Println(table.Render())
}
