package main

import (
	"strings"
)

func cleanInput(test string) []string {
	items := strings.Fields(test)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	// fmt.Printf("ITEMS: %#v\n", items)
	return items
}
