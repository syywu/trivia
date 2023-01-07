package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Input(s string, r *bufio.Reader) (string, error) {
	fmt.Print(s)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func AskName() {
	name, _ := Input("What is your name?\n", bufio.NewReader(os.Stdin))
	fmt.Println("Welcome to Trivia,", name)
	time.Sleep(2 * time.Second)
	Q1()
}

func Q1() {
	q, _ := Input("In a website browser address bar, what does “www” stand for?\na-Word Wide West\nb-World Wild Web\nc-World Wide Web\nd-World Wide Window\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Correct")
		time.Sleep(2 * time.Second)
		Q2()
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q1()
	}
}

func Q2() {
	q, _ := Input("In a bingo game, which number is represented by the phrase “two little ducks”?\na-22\nb-88\nc-55\nd-66\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Correct")
		time.Sleep(2 * time.Second)
		Q3()
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q2()
	}
}

func Q3() {
	q, _ := Input("According to Greek mythology, who was the first woman on earth?\na-Eve\nb-Athena\nc-Socrates\nd-Pandora\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Correct")
		time.Sleep(2 * time.Second)
		Q4()
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q3()
	}
}

func Q4() {
	q, _ := Input("Which bone are babies born without?\na-Patella\nb-Fibula\nc-Clavicles\nd-Tibia\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q5()
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q4()
	}
}

func Q5() {
	q, _ := Input("How many hearts does an octopus have?\na-4\nb-3\nc-2\nd-8\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q6()
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q5()
	}
}

func Q6() {
	q, _ := Input("What type of music has been shown to help plants grow better and faster?\na-Hip Hop\nb-Deep House\nc-Classical\nd-Lo-Fi\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q7()
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q6()
	}
}

func Q7() {
	q, _ := Input("What’s the hardest rock?\na-Granite\nb-Diamond\nc-Crystal\nd-Iron\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q8()
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q7()
	}
}

func Q8() {
	q, _ := Input("What country has the world’s most ancient forest?\na-Australia\nb-Colombia\nc-Brazil\nd-Kenya\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q9()
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Wrong answer")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q8()
	}
}

func Q9() {
	q, _ := Input(" Which fruit floats because 25% of its volume is air?\na-Orange\nb-Grapefruit\nc-Watermelon\nd-Apple\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		Q10()
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q9()
	}
}

func Q10() {

	q, _ := Input("When did the Cold War end?\na-1985\nb-1980\nc-1990\nd-1989\n", bufio.NewReader(os.Stdin))
	switch q {
	case "a":
		fmt.Println("Wrong answer")
	case "b":
		fmt.Println("Wrong answer")
	case "c":
		fmt.Println("Wrong answer")
	case "d":
		fmt.Println("Correct")
		time.Sleep(time.Second)
		fmt.Println("Congrats! You have one big brain :D")
	default:
		fmt.Println("Invalid option")
		time.Sleep(time.Second)
		Q10()
	}
}

func main() {
	AskName()
}
