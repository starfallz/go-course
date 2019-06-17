package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, strings.Join(sortLetters(letters("aba")), "\n"))
}

func letters(s string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, currentLetter := range s {
		letterCount[currentLetter]++
	}

	return letterCount
}

func sortLetters(m map[rune]int) []string {
	var allKeys []rune

	for key, _ := range m {
		allKeys = append(allKeys, key)
	}

	sortRune(allKeys)

	var result []string

	for _, currentKey := range allKeys {
		result = append(result, strings.Join([]string{string(currentKey), strconv.Itoa(m[currentKey])}, ":"))
	}

	return result
}

func sortRune(r []rune) {
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
}
