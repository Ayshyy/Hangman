package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	isStart bool
	words   string
)

func main() {
	load()
	Init()
}

func load() {
	//Load File words.txt in memory
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words = string(content)

}

func selectword(line int) string {
	for u, e := range strings.Split(words, "\n") {
		if u == line-1 {
			return e
		}
	}
	return ""
}

func Init() {
	input := ""
	fmt.Println("Appuyez sur entré pour commencé !")
	fmt.Scan(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}
