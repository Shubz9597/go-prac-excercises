package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "172.25.97.209"
	port     = 5432
	user     = "postgres"
	password = "rootroot"
	dbname   = "phoneNumbers"
)

type DB struct {
	db *sql.DB
}

type PhoneNumber struct {
	id     int
	number string
}

func ConnectDb() (*DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("there is some error connecting with database %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("there is some error pinging the database %w", err)
	}

	var database DB

	database.db = db

	return &database, nil
}

func (db *DB) CloseDb() {
	db.db.Close()
}

func (db *DB) InsertRecords() {
	sqlStatement := `INSERT INTO public."phone" (number) 
VALUES 
    ('1234567890'), 
    ('123 456 7891'), 
    ('(123) 456 7892'), 
    ('(123) 456-7893'), 
    ('123-456-7894'), 
    ('123-456-7890'), 
    ('1234567892'), 
    ('(123)456-7892')`

	_, err := db.db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func (db *DB) CheckRecords() []PhoneNumber {
	sqlStatement := `SELECT * from public."phone"`

	data, err := db.db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer data.Close()
	var phone []PhoneNumber
	for data.Next() {

		var id int
		var number string
		err := data.Scan(&id, &number)
		if err != nil {
			panic(err)
		}

		phone = append(phone, PhoneNumber{id: id, number: number})
	}

	return phone
}

func updatePhoneNumber(number string) string {
	normalize := strings.ReplaceAll(number, " ", "")
	normalize = strings.ReplaceAll(normalize, "(", "")
	normalize = strings.ReplaceAll(normalize, ")", "")
	normalize = strings.ReplaceAll(normalize, "-", "")
	fmt.Println(normalize)
	return normalize
}

func (db *DB) UpdateAndDeleteRecordsAfterNormalizing(phoneNumbers []PhoneNumber) {
	tx, err := db.db.Begin()

	if err != nil {
		log.Fatal("Error starting transaction %w", err)
	}

	updateStmt := `
		UPDATE public."phone"
		SET number = $1
		WHERE id = $2
	`
	for _, val := range phoneNumbers {
		_, err := tx.Exec(updateStmt, updatePhoneNumber(val.number), val.id)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error updating phone numbers: %v", err)
		}
	}

	deleteStmt := `
        WITH duplicates AS (
            SELECT id,
                ROW_NUMBER() OVER (PARTITION BY number ORDER BY id) AS rn
            FROM public."phone"
        )
        DELETE FROM public."phone"
        WHERE id IN (
            SELECT id FROM duplicates WHERE rn > 1
        )
    `

	_, err = tx.Exec(deleteStmt)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error deleting duplicates: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}

	fmt.Println("Records have been Normalized")
}
