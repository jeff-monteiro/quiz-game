package main

import (
	"bufio"
	"encoding/csv"
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
	Points    string
	Questions []Question
}

// Method that gives the kick off on game
func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome:")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name
	fmt.Printf("Vamos jogar %s", g.Name)
}

// Method that open, process and read the CSV file
func (g *GameState) ProccessCSV() {

	f, err := os.Open("Questionsgo.csv")
	if err != nil {
		panic("Erro ao abrir arquivo CSV!")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler arquivo CSV!")
	}

	for index, record := range records {
		fmt.Println(index, record)
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}

	}

}

func main() {
	game := &GameState{}
	go game.ProccessCSV()
	game.Init()

	fmt.Println(game)
}

// Function to make conversion between string and int types
func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Erro na conversao de tipos")
	}

	return i
}
