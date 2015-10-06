package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	BookStore      []Book   = make([]Book, 0)
	Authors        []Author = make([]Author, 0)
	ErrAuthorEmpty          = errors.New("Author can not be empty")
	path                    = filepath.Join("/tmp", "bookstore")
	filename                = "list"
)

type Author struct {
	ID          int
	Name        string
	Nationality string
}

type Book struct {
	ISBN   string
	Author *Author
	Title  string
	Pages  int
}

type Message struct {
	Error string
}

func saveBook(b Book) {
	BookStore = append(BookStore, b)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	file := filepath.Join(path, filename)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	marshaled, err := json.MarshalIndent(BookStore, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(file, marshaled, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(marshaled))
}

func index(c *echo.Context) error {
	if len(BookStore) == 0 {
		return nil
	}
	c.JSON(201, BookStore)
	return nil
}

func addBook(c *echo.Context) error {
	var b Book
	b.Author = nil

	err := c.Bind(&b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if b.Author == nil {
		var m Message
		m.Error = "Author can not be empty."
		c.JSON(203, m)
		return ErrAuthorEmpty
	}

	err = c.JSON(201, b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Added ", b.Title)
	saveBook(b)
	return nil
}

func main() {
	e := echo.New()
	e.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/book", 301)
	})

	// Middleware
	e.Use(middleware.Logger())

	// Handlers
	e.Get("/book", index)
	e.Post("/book", addBook)

	fmt.Println("Listening on port :5000 ")
	e.Run(":5000")
}