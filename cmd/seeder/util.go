package main

import (
	"encoding/json"
	"log"
	"os"
)

func readFileToJson(filePath string, data any) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, data)
	if err != nil {
		log.Fatal(err)
	}
}
