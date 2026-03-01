package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"time"
	"math/rand"
	"os"
)

type Flashcard struct {
	Question string
	Answer   string
}

type Pair struct {
	Q string
	A string
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


func randomSelection()  (string, string) {
	var pairs []Pair

	records := openCsvFile("test.csv")
	rand.Seed(time.Now().Unix())


	for i, record := range records {
		if i == 0{
			continue
		}

		if len(record) < 2{
			continue
		}

		pairs = append(pairs, Pair {
			Q: record[0],
			A: record[1],
		})
	}

	rand.Seed(time.Now().Unix())
	randonIndex := rand.Intn(len(pairs))
	randomPair := pairs[randonIndex]

	question := fmt.Sprint(randomPair.Q) 
	answer :=  fmt.Sprint(randomPair.A)
	return question,answer

}

func main() {
	fmt.Println("Question # <WIP> lol:")
	q, a := randomSelection()
	fmt.Println(q)
	fmt.Println(a)

 }
