package main

import (
	"fmt"
	"log"
	"phoneNumberNormalizer/model"
)

func main() {
	fmt.Println("nigga money")

	//Connect with the database
	db, err := model.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	phoneNumbers := db.CheckRecords()

	if len(phoneNumbers) == 0 {
		db.InsertRecords()
	} else {
		db.UpdateAndDeleteRecordsAfterNormalizing(phoneNumbers)
	}

	defer db.CloseDb()
}
