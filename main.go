package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

// Method that gives the kick off on game
func (g *GameState) Init() {
	fmt.Println("Welcome to the quiz")
	fmt.Println("Choose a nickname:")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name
	fmt.Printf("Let's Play %s", g.Name)
}

// Method that open, process and read the CSV file
func (g *GameState) ProccessCSV() {

	f, err := os.Open("Questionsgo.csv")
	if err != nil {
		panic("Error on open CSV file!")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error on read CSV file!")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}

	}

}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[35m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Insert one alternative:")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}
		fmt.Println(answer)

		if answer == question.Answer {
			fmt.Println("Congratulations, Right answer!")
			g.Points += 10
		} else {
			fmt.Println("Ops! You were wrong!")
			fmt.Println("<><><><><><><><><><><><><><><><><><><>")
		}
	}

}

func main() {
	game := &GameState{}
	go game.ProccessCSV()
	game.Init()
	game.Run()

	fmt.Printf("End Game, you performed %d points!\n", game.Points)

}

// Function to make conversion between string and int types
func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("Not allowed insert charactere different of a number!")
	}

	return i, nil
}
