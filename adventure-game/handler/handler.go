package handler

import (
	"adventureGame/types"
	"encoding/json"
	"fmt"
	"os"
)

var ParsedData map[string]types.Stories

func ParseJson() (map[string]types.Stories, error) {
	fileName := "gopher.json"

	reader, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("there is some error reading the file %w", err)
	}

	err = json.Unmarshal(reader, &ParsedData)

	if err != nil {
		return nil, fmt.Errorf("there is some problem Unmarshaling Json data %w", err)
	}

	// fmt.Println(parsedData)
	return ParsedData, nil
}
