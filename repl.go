package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			inputString := scanner.Text()
			tokens := cleanInput(inputString)
			if len(tokens) > 0 {
				fmt.Printf("Your command was: %s\n", tokens[0])
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

func cleanInput(test string) []string {
	items := strings.Fields(test)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	// fmt.Printf("ITEMS: %#v\n", items)
	return items
}
