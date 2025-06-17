// Problem: Word Counter
// Write a program that counts the number of words in a given string. Words are separated by spaces, and you should ignore extra whitespace.

package main

import (
	"fmt"
	"strings"
)

func countWords(text string) int {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return 0
	}
	words := strings.Fields(trimmed)
	return len(words)
}

func main() {

	testCases := []string{
		"Hello world from Go",
		"  Go   is   awesome  ",
		"",
		"   ",
		"SingleWord",
		"One Two Three Four Five",
	}

	fmt.Println("Word Counter Program")
	fmt.Println("===================")

	for i, text := range testCases {
		count := countWords(text)
		fmt.Printf("Test %d: \"%s\" -> %d words\n", i+1, text, count)
	}

	fmt.Println("\nTry your own text:")
	var userInput string
	fmt.Print("Enter a sentence: ")

	fmt.Scanln(&userInput)

	if userInput != "" {
		result := countWords(userInput)
		fmt.Printf("Your sentence has %d words.\n", result)
	}
}
