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
	fmt.Println("Welcome to Trivia,", name)
	Q1()
}

func Q1() {
	q, _ := Input("In a website browser address bar, what does “www” stand for?\na- Word Wide West\nb- World Wild Web\nc- World Wide Web\nd-World Wide Window\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Correct")
		Q2()
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		Q1()
	}
}

func Q2() {
	q, _ := Input("In a bingo game, which number is represented by the phrase “two little ducks”?\na-22\nb-88\nc-55\nd-66\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Correct")
		Q3()
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		Q2()
	}
}

func Q3() {
	q, _ := Input("According to Greek mythology, who was the first woman on earth?\na- Eve\nb- Athena\nc- Socrates\n d- Pandora\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Correct")
	default:
		fmt.Println("Invalid option")
		Q3()
	}
}

func main() {
	AskName()
}
