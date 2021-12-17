package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func ConnectToDB() *sql.DB {
	if !doesDbFileExist() {
		log.Println("Creating db.sqlite3...")
		file, err := os.Create("db.sqlite3") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("db.sqlite3 created")
	}

	db, err := sql.Open("sqlite3", "./db.sqlite3") // Open the created SQLite File

	if err != nil {
		panic(err)
	}

	log.Println("Connected to DB")
	return db
}

func DoesTableExist(tableName string) bool {
	_, tableCheck := DB.Query("select * from " + tableName + ";")

	return tableCheck == nil

}

func doesDbFileExist() bool {
	if _, err := os.Stat("./db.sqlite3"); err == nil {
		// path/to/whatever exists
		return false
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Println("File may or may not exist... I can't find it :(")
		return true
	}
}

func CreateUrlTable() {
	createUrlTableSQL := `CREATE TABLE url (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"original_url" TEXT,
		"short_url" TEXT		
  	);` // SQL Statement for Create Table

	log.Println("Creating url table...")
	statement, err := DB.Prepare(createUrlTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println("Failed to create url table!")
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("url table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertUrl(oUrl string, sUrl string) {
	log.Println("Inserting url record ...")
	insertUrlSQL := `INSERT INTO url(original_url, short_url) VALUES ($1, $2)`

	_, err := DB.Exec(insertUrlSQL, oUrl, sUrl)
	if err != nil {
		panic(err)
		//log.Fatalln(err.Error())
	}
}

type UrlData struct {
	ID          int64
	ShortUrl    string
	OriginalUrl string
}

func GetUrlFromDB(sUrl string) string {
	retrieveUrlSQL := `SELECT * FROM url WHERE short_url = $1`
	var data UrlData

	row := DB.QueryRow(retrieveUrlSQL, sUrl)
	err := row.Scan(&data.ID, &data.OriginalUrl, &data.ShortUrl)

	switch err {
	case sql.ErrNoRows:
		return "NO URL FOUND"

	case nil:
		return data.OriginalUrl

	default:
		panic(err)
	}

	return "SOMETHING WENT WRONG"
}
