package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"time"
	"math/rand"
	"os"
)

var (
 Reset = "\033[0m"
 Red = "\033[31m"
 Green = "\033[32m"
 Yellow = "\033[33m"
 Blue = "\033[34m"
 Magenta = "\033[35m"
 Cyan = "\033[36m"
 Gray = "\033[37m"
 White = "\033[97m"
)


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
	//TODO: Implement quesiton count
	fmt.Println("Question # $questionNumber:\n\n")
	q, a := randomSelection()
	fmt.Println("Question:", q)
	
	//There has to be a better way to implement this. I haven't had good experiences with Scanln
	fmt.Println(Yellow + "\n\nType the command or answer and then press enter to reveal the answer..." + Reset)
	fmt.Print("Command -> $ ")
	fmt.Scanln()
	fmt.Println(Green + a + Reset)
 }
