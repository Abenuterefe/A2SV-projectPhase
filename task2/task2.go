package main

import (
	"fmt"
	"strings"
	"unicode"
)

// CountWords returns a map of word frequencies from the input string.
func CountWords(text string) map[string]int {
	if len(text) == 0 {
		return map[string]int{}
	}

	// Normalize input: lowercase and remove punctuation
	normalized := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, text)

	wordList := strings.Fields(normalized)

	wordCount := make(map[string]int)
	for _, w := range wordList {
		wordCount[w]++
	}

	return wordCount
}

// CheckPalindrome returns true if input is a palindrome (ignores case and punctuation)
func CheckPalindrome(s string) bool {
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	for left, right := 0, len(filtered)-1; left < right; left, right = left+1, right-1 {
		if filtered[left] != filtered[right] {
			return false
		}
	}
	return true
}

func main() {
	sampleText := "Hello, world! Hello, universe. Hello world!"
	wordStats := CountWords(sampleText)

	fmt.Println("Word Frequencies:")
	for word, count := range wordStats {
		fmt.Printf("  %s: %d\n", word, count)
	}

	fmt.Println("\nPalindrome Checks:")
	fmt.Println("Is 'A man, a plan, a canal: Panama' a palindrome?", CheckPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("Is 'Hello, world!' a palindrome?", CheckPalindrome("Hello, world!"))
}
