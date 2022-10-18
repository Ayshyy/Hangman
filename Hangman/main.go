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
	// Debug() // temp fonction delete after finish
	Play()
}

func Debug() {
	fmt.Println(hangmanData.ToFind)
	fmt.Println(CheckWin())
	fmt.Println(hangmanData.HangmanPositions)
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

	words = ""

	//Load Hangman.txt
	hangman, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	HangmanPositions = string(hangman)

	//set Word hide
	hangmanData.Word = HideWord(hangmanData.ToFind)
	hangmanData.Word = AddRandomLetter()
	//hangmanData.Word = AddRandomLetter()
	hangmanData.Attempts = 9
}

func finish() {
	fmt.Println("Game finish, press enter to close\n \n ")
	fmt.Println("Tape Y pour rejouer !")
	scanner := bufio.NewReader(os.Stdin)
	char, _, _ := scanner.ReadRune()
	switch char {
	case 'y':
		main()
	default:
		os.Exit(1)
	}
}

func SelectWord(line int) string {
	for u, e := range strings.Split(words, "\n") {
		if u == line-1 {
			return strings.ReplaceAll(e, " ", "")
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
	for i := 0; i < len(word)-1; i++ {
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

func AddRandomLetter() string {
	result := ""
	rdm := getRandomNumber(len(hangmanData.ToFind)-1, 0)
	strConvert := []rune(hangmanData.Word)
	for i, r := range hangmanData.ToFind {
		letter := ' '
		if i == rdm {
			letter = r
			strConvert[i] = r
		}
		if r == letter && letter != ' ' {
			strConvert[i] = r
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

func CheckWin() bool {
	strConv := []rune(hangmanData.Word)
	strWord := []rune(hangmanData.ToFind)
	nbrValid := 0
	WordSize := 0

	for i := 0; i < len(hangmanData.ToFind)-1; i++ {
		if strConv[i] == strWord[i] && strWord[i] != ' ' && strWord[i] != '\n' {
			nbrValid++
		}
	}
	for range hangmanData.Word {
		WordSize++
	}

	return WordSize == nbrValid
}

func getHangmanFromPos(position int) string {
	position++
	HangmanSize := 8
	if position > 10 && position < 0 {
		fmt.Printf("Hangman position Error: %v", position)
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
	fmt.Println("Tentative Restante: " + strconv.Itoa(hangmanData.Attempts))
	fmt.Print("Tape une lettre : ")
	for scanner.Scan() && isStart {
		switch scanner.Text() {
		case "finish":
			finish()
		default:
			if len(scanner.Text()) == 1 {
				if hangmanData.Attempts >= 1 {

					if CheckRune(TextToRune(scanner.Text())) {
						fmt.Println("Lettre trouvée: " + scanner.Text())
						hangmanData.Word = addFindLetter(TextToRune(scanner.Text()))
						fmt.Println(hangmanData.Word)
						if CheckWin() {
							fmt.Println("tu as gagné !")
							finish()
						}
					} else {
						fmt.Println(getHangmanFromPos(hangmanData.HangmanPositions))
						hangmanData.Attempts -= 1
						hangmanData.HangmanPositions += 1
						fmt.Println(hangmanData.Word)
						fmt.Println("Tentative Restante: " + strconv.Itoa(hangmanData.Attempts))
					}
				} else {
					fmt.Println(getHangmanFromPos(hangmanData.HangmanPositions))
					fmt.Println("Pendu")
					finish()
				}
			} else if len(scanner.Text()) == 0 {

			} else {
				fmt.Println("Trop de lettres.")
			}
		}
		fmt.Print("Tape une lettre : ")
	}
}
