package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Flashcard struct {
	Question string
	Answer   string
}


func openCsvFile(csvPath string) [][]string {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal("Unable to open file..." + csvPath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse the file as csv..." + csvPath, err)
	}
	return records
}

func main() {
	records := openCsvFile("test.csv")

		for _, row := range records {
		if len(row) >= 2 {
			fmt.Printf("Q: %s\nA: %s\n\n", row[0], row[1])
		}
	}
}
