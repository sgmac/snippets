package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	words := []string{
		"anna",
		"civic",
		"kAyak",
		"level",
	}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("word %s is palindrome\n", word)
		}
	}

}

func isPalindrome(word string) bool {
	var once sync.Once
	lower := func() {
		word = strings.ToLower(word)
	}
	once.Do(lower)

	word = strings.ToLower(word)
	end := len(word) - 1
	if end == 0 {
		return true
	}
	if end == 1 && (word[0] == word[end]) {
		return true
	}

	if word[0] != word[end] {
		return false
	}
	return isPalindrome(string(word[1:end]))
}
