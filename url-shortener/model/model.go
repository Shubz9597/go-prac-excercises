package model

import (
	"database/sql/driver"
	"log"

	"github.com/lib/pq"
)

func Connect() driver.Conn {
	connString := "user=postgres dbname=parsedUrls password=rootroot sslmode=disable"

	db, err := pq.Open(connString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
