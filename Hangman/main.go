package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	isStart bool
)

func main() {
	Init()

}
func load() {
	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
func selectword(i int){
	for i, r := range strings.Split(words, "\n") {
	if u == line(
		return e
	)
	}
	return ""
}

func Init() {
	input := ""
	fmt.Println("Appuyez sur une touche pour commencé !")
	fmt.Scanf(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}
