package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(s string, r *bufio.Reader) (string, error) {
	fmt.Print(s)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func AskName() {
	name, _ := Input("What is your name?", bufio.NewReader(os.Stdin))
	fmt.Println("Welcome to Trivia", name)
}

// Trivia Question: In a website browser address bar, what does “www” stand for?
// Answer: World Wide Web
func Q1() {
	q, _ := Input("In a website browser address bar, what does “www” stand for?\na- Word Wide West\nb- World Wild Web\nc- World Wide Web\nd-World Wide Window\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Correct")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		Q1()
	}
}

func main() {
	AskName()
	Q1()
}
