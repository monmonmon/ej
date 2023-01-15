package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func getWordFromArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New(":P")
	}
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		if len(word) >= 2 {
			return word, nil
		}
	}
	return "", errors.New(":P")
}

func main() {
	word, err := getWordFromArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	databasefilepath := filepath.Join(homeDir, "lib/ejdict.sqlite3")
	if _, err := os.Stat(databasefilepath); err != nil {
		println("dictionary file does not exist")
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", databasefilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select word, mean from items where word like ?", word+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var word string
		var mean string
		err = rows.Scan(&word, &mean)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(word)
		lines := strings.Split(mean, " / ")
		for _, line := range lines {
			fmt.Println("\t", line)
		}
	}
}
