package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	isStart     bool
	words       string
	hangmanData HangManData
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	load()
	//Init()
	Debug() // temp fonction delete after finish
}

func Debug() {
	fmt.Println(hangmanData.ToFind)
}

func load() {
	//Load File words.txt in memory
	fmt.Println("Hangman is Loading...")
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words = string(content)
	//Init hangman Data
	//Select random word
	hangmanData.ToFind = selectword(getRandomNumber(WordsSize(), 0))
}

func selectword(line int) string {
	for u, e := range strings.Split(words, "\n") {
		if u == line-1 {
			return e
		}
	}
	return ""
}

func getRandomNumber(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result
}

func WordsSize() int {
	result := 0
	for range strings.Split(words, "\n") {
		result++
	}
	return result
}

func Init() {
	input := ""
	fmt.Println("Appuyez sur entré pour commencé !")
	fmt.Scan(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}
