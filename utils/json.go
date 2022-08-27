package utils

import (
	"encoding/json"
	"log"
)

func Print(result any) {
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", jsonData)
}
