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
	HangmanPositions int    // hangman position state
}

func main() {
	load()
	// Debug() // temp fonction delete after finish
	Play()
}

func load() { //Chargement des fichiers du jeux
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

	fmt.Println(getHangmanFromPos(9)) //Affichage du hangma
	fmt.Println("\n\n\n ----------------")
	fmt.Println("  ")
}

func finish() { //Met fin au jeux. Possibilité de rejouer en tapant sur Y autrement le programme s'arrête.
	isStart = false
	getHangmanFromPos(hangmanData.HangmanPositions)
	fmt.Println("\n\nGame finish, press enter to close\n ")
	fmt.Println("Tape Y pour rejouer !")
	scanner := bufio.NewReader(os.Stdin)
	char, _, _ := scanner.ReadRune()
	switch char {
	case 'y':
		main()
	case 'Y':
		main()
	default:
		os.Exit(1)
	}
}

func SelectWord(line int) string { //On récupère un mot selon la ligne où il se situe dans le fichier word.txt
	for index, str := range strings.Split(words, "\n") {
		if index == line-1 {
			return strings.ReplaceAll(str, " ", "")
		}
	}
	return ""
}

func CheckRune(rune rune) bool { //On verifie si la rune correspond à une lettre dans le mot à trouver
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

func addFindLetter(runes rune) string { //on ajoute
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

func AddRandomLetter() string { //On ajoute une lettre aléatoire
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

func WordsSize() int { //Retourne le nombre de mot contenu dans le fichier word.txt
	result := 0
	for range strings.Split(words, "\n") {
		result++
	}
	return result
}

func getRandomNumber(max, min int) int { //on récupère un nombre aléatoire compris entre min et max
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result
}

func CheckWin() bool { //On verifie si le mot trouvé et le mot rechercher correspondes dans ce cas il retourne true ou false
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

func checkWord(word string) bool { //détermine si le mot taper et valide, si le mots contient moin de lettre c'est considéré comme mal tapé.
	strConv := []rune(word)
	strWord := []rune(hangmanData.ToFind)
	nbrValid := 0
	WordSize := 0
	if len(word) != len(strConv) {
		return false
	}

	for i := 0; i < len(hangmanData.ToFind)-1; i++ {
		if strConv[i] == strWord[i] && strWord[i] != ' ' && strWord[i] != '\n' {
			nbrValid++
		}
	}
	for range hangmanData.Word {
		WordSize++
	}

	valid := WordSize == nbrValid //Si le mot tapé est incorrect on retire 2 essais et on affiche l'état du pendu
	if !valid {
		punishment := 2 // determine le nombre de point retiré au joueur
		hangmanData.Attempts -= punishment
		hangmanData.HangmanPositions += punishment
		getHangmanFromPos(hangmanData.HangmanPositions)
	}
	return valid
}

func getHangmanFromPos(position int) string {
	position++
	HangmanSize := 8
	if position > 10 && position < 0 {
		fmt.Printf("Hangman position Error: %v", position)
		return ""
	}
	hangmanStat := ""
	for index, s := range strings.Split(HangmanPositions, "\n") {
		INC := HangmanSize * position
		if index >= INC-HangmanSize-1 && index < INC-1 {
			hangmanStat += s + "\n"
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
						fmt.Println(">> Lettre trouvée: " + scanner.Text() + "\n")
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
					fmt.Println("               Pendu !")
					fmt.Println("\nLe mot était : " + hangmanData.ToFind + "")
					finish()
				}
			} else if len(scanner.Text()) == 0 {

			} else {

				if checkWord(scanner.Text()) {
					fmt.Println("\n\nBravo ! Le mot était bien: " + hangmanData.ToFind)
					finish()
				}
				fmt.Println("Trop de lettres.")
			}
		}
		fmt.Print("Tape une lettre : ")
	}
}
