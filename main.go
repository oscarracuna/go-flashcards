package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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
	fmt.Println("test")
	records := openCsvFile("deck.csv")
	fmt.Println(records)
}
