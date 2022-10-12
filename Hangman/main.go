package main

import "fmt"

var (
	isStart bool = false
)

func main() {
	Init()
}

func Init() {
	input := ""
	fmt.Println("Appuyez sur une touche pour commencé !")
	fmt.Scanf(">", &input)
	for len(input) > 1 {
	}
	isStart = true
}
