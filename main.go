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
	fmt.Println("Hi", name)
}

func main() {
	AskName()
}
