package csvRead

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(fileName string) [][]string {
	var fileN = fileName + ".csv"
	file, err := os.Open(fileN)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return data
}
