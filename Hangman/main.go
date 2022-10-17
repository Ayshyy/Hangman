package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	words            string
	isStart          bool   //Check if party is Running
	HangmanPositions string //Data from file hangman.txt and word.txt
	hangmanData      HangManData
)

type HangManData struct {
	Word             string // Word composed of '_', ex: H_ll_
	ToFind           string // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int    // Number of attempts left
	HangmanPositions int    // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	load()
	//Debug() // temp fonction delete after finish
	Play()
}

func Debug() {
	fmt.Println(hangmanData.ToFind)
	fmt.Println("--------Debug-------")
	for i := 0; i <= 10; i++ {
		fmt.Println(getHangmanFromPos(i))
	}
	fmt.Println("--------------------")
}

func load() {
	//Load words.txt
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

	//set Word hide
	hangmanData.Word = HideWord(hangmanData.ToFind)
	hangmanData.Attempts = 10
}
func finish() {
	fmt.Println("Game finish, Good bye")
	os.Exit(1)
}

func SelectWord(line int) string {
	for u, e := range strings.Split(words, "\n") {
		if u == line-1 {
			return e
		}
	}
	return ""
}

func CheckRune(rune rune) bool {
	for _, r := range hangmanData.ToFind {
		if rune == r {
			return true
		}
	}
	return false
}

func HideWord(word string) string {
	str := ""
	for range word {
		str += "_"
	}
	return str
}
func addFindLetter(runes rune) string {
	result := ""
	strConvert := []rune(hangmanData.Word)
	for i, r := range hangmanData.ToFind {
		if r == runes {
			strConvert[i] = runes
		}
	}
	for _, r := range strConvert {
		result += string(r)
	}
	return result
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
func TextToRune(text string) rune {
	for _, r := range text {
		return r
	}
	return ' '
}
func Play() {
	isStart = true
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(hangmanData.Word)
	fmt.Print("Tape une lettre : ")
	for scanner.Scan() && isStart {
		switch scanner.Text() {
		case "finish":
			finish()
		default:
			if len(scanner.Text()) == 1 {
				if hangmanData.Attempts >= 1 {
					fmt.Println("Tentative Restante: " + strconv.Itoa(hangmanData.Attempts))
					if CheckRune(TextToRune(scanner.Text())) {
						fmt.Println("Lettre trouvée: " + scanner.Text())
						hangmanData.Word = addFindLetter(TextToRune(scanner.Text()))
						fmt.Println(hangmanData.Word)
					} else {
						hangmanData.Attempts -= 1
						hangmanData.HangmanPositions += 1

						fmt.Println(getHangmanFromPos(hangmanData.HangmanPositions))
						fmt.Println(hangmanData.Word)
					}
				} else {
					fmt.Println("Pendu")
					finish()
				}
			} else {
				fmt.Println("Trop de lettres.")
			}
		}
		fmt.Print("Tape une lettre : ")
	}
}
