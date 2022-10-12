package main

import "fmt"

var (
	isStart bool
	words   []string
)

func main() {
	Init()

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
