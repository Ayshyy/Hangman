package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	isStart bool
	words   []string
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

func load() {

}

func Init() {
	input := ""
	fmt.Println("Appuyez sur entré pour commencé !")
	fmt.Scan(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}
