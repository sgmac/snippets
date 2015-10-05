package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	BookStore []Book   = make([]Book, 0)
	Authors   []Author = make([]Author, 0)
)

type Author struct {
	ID          int
	Name        string
	Nationality string
}

type Book struct {
	ISBN   string
	Author Author
	Title  string
	Pages  int
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
	var a Author
	var b Book
	b.Author = a

	c.Bind(&b)
	c.JSON(201, b)
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
