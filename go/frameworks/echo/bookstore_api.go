package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	BookStore      []Book   = make([]Book, 0)
	Authors        []Author = make([]Author, 0)
	ErrAuthorEmpty          = errors.New("Author can not be empty")
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

func index(c *echo.Context) error {
	if len(BookStore) == 0 {
		return nil
	}
	for _, b := range BookStore {
		c.JSON(201, b)
	}
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

	BookStore = append(BookStore, b)
	fmt.Println("Added ", b.Title)
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
