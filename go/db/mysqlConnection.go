package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db       string
	host     string
	password string
	port     string
	user     string
)

func init() {
	flag.StringVar(&db, "d", "test", "DBname")
	flag.StringVar(&host, "h", "localhost", "Hostname")
	flag.StringVar(&password, "p", "", "Password")
	flag.StringVar(&port, "P", "3306", "Port")
	flag.StringVar(&user, "u", "root", "User")
}

func main() {
	flag.Parse()
	var connectString string
	var value string = "hello world"

	if flag.NArg() >= 1 {
		value = strings.Join(flag.Args(), " ")
	}

	if password != "" {
		connectString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	} else {
		connectString = fmt.Sprintf("%s@tcp(%s:%s)/%s", user, host, port, db)
	}
	db, err := sql.Open("mysql", connectString)

	// Define at least one connection idle on the pool
	db.SetMaxIdleConns(1)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Continue...")

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS test")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil {
		log.Fatal(err)
	}
	sql := fmt.Sprintf("INSERT INTO test.hello(world) VALUES('%s')", value)
	// res, err := db.Exec("INSERT INTO test.hello(world) VALUES('hello world!')")
	res, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted  %d rows", rowCount)

	rows, err := db.Query("SELECT * from test.hello")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reading just one row")
	var oneRow string
	err = db.QueryRow("SELECT * from test.hello limit 1").Scan(&oneRow)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("one Row: ", oneRow)

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Cols: ", cols)

	log.Println("Working with transactions")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Query("UPDATE hello SET world = 'NEW WORLD!'")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %q", s)
	}
	res, err = db.Exec("DELETE FROM test.hello LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("deleted rows:", rowCnt)

	res, err = db.Exec(
		"INSERT INTO test.hello(world) VALUES(?)", "hello world!")

	rows.Close()
}
