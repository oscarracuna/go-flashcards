package main

import (
	"encoding/csv"
	"bufio"
	"fmt"
	"log"
	"time"
	"math/rand"
	"os"
	"os/exec"
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

func getCsv(csvPath string) [][]string {
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

func getFlashcards()  (string, string) {
	var pairs []Pair

	records := getCsv("test.csv")
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

func displayQuestion() {
	scanner := bufio.NewScanner(os.Stdin)

	counter := 1
	i := 1
	for i == 1 {
		fmt.Printf(Yellow+ "Question #%v:\n" + Reset, counter)
		q, a := getFlashcards()
		fmt.Println(q)
		
		fmt.Println(Yellow + "\nType the command or answer and then press enter to reveal the answer..." + Reset)
		fmt.Print("-> $ ")
		scanner.Scan()
		
		fmt.Println(Green + "\nAnswer:" + Reset)
		fmt.Println(a)
		
		fmt.Println(Yellow + "\nPress enter to go to the next question or press ctrl+c to exit." + Reset)
		scanner.Scan()
		clearScreen()
		counter += 1
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func main() {
	fmt.Print(Green + "Welcome back!\nEnjoy your study session.\n\n" + Reset)
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
	clearScreen()
	displayQuestion()
}
