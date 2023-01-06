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

func main() {
	Input("What is your name?", bufio.NewReader(os.Stdin))
}
