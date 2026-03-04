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
	"strings"
	tea "charm.land/bubbletea/v2"
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


//==================
// Start of UI thing
//==================
var choices = []string{"Option1", "Option2", "option3"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := strings.Builder{}
	s.WriteString("Here you have a couple of options.\n\n")

	for i := range choices {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return tea.NewView(s.String())
}

func displayThing() {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}
}


//=================
// End of UI thing
//=================




type Pair struct {
	Q string
	A string
}

func getCsv(csvPath string) [][]string {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal("Unable to open file -> " + csvPath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse the file as csv -> " + csvPath, err)
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
	displayThing()
	fmt.Print(Green + "Welcome back!\nEnjoy your study session.\n\n" + Reset)
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
	clearScreen()
	displayQuestion()
}
