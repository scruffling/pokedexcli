package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	startRepl()
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			inputString := scanner.Text()
			tokens := cleanInput(inputString)
			if len(tokens) > 0 {
				fmt.Printf("Your command was: %v\n", tokens[0])
			}
		} else {
			// check for errors if Scan() is false
			if err := scanner.Err(); err != nil {
				fmt.Printf("You have an error: %v\n", err)
			}
			// false means EOF or error; exit the loop
			break
		}
	}
}
