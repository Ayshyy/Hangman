package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	isStart                 bool   //Check if party is Running
	words, HangmanPositions string //Data from file hangman.txt and word.txt
	hangmanData             HangManData
)

type HangManData struct {
	Word             string // Word composed of '_', ex: H_ll_
	ToFind           string // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int    // Number of attempts left
	HangmanPositions int    // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	load()
	//Init()
	Debug() // temp fonction delete after finish
}

func Debug() {
	fmt.Println(hangmanData.ToFind)
	fmt.Println("---------------")
	for i := 0; i <= 10; i++ {
		fmt.Println(getHangmanFromPos(i))
	}
	fmt.Println("---------------")
}

func load() {
	//Load File words.txt in memory
	fmt.Println("Hangman is Loading...")
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Init hangman Data
	words = string(content)

	//Select random word
	hangmanData.ToFind = SelectWord(getRandomNumber(WordsSize(), 0))

	//Load Hangman.txt
	hangman, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	HangmanPositions = string(hangman)
}

func SelectWord(line int) string {
	for u, e := range strings.Split(words, "\n") {
		if u == line-1 {
			return e
		}
	}
	return ""
}

func WordsSize() int {
	result := 0
	for range strings.Split(words, "\n") {
		result++
	}
	return result
}

func getRandomNumber(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result
}

func getHangmanFromPos(position int) string {
	HangmanSize := 8
	if position > 10 && position < 1 {
		fmt.Printf("Hangman position Error: %v", position)
		os.Exit(3)
		return ""
	}
	hangmanStat := ""
	for index, r := range strings.Split(HangmanPositions, "\n") {
		INC := HangmanSize * position
		if index >= INC-HangmanSize-1 && index < INC-1 {
			hangmanStat += r + "\n"
		}
	}
	return hangmanStat
}

func Init() {
	input := ""
	fmt.Println("Appuyez sur entré pour commencé !")
	fmt.Scan(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}

func Play() {
	for isStart {

	}
}
