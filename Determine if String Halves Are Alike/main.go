package main

import "fmt"

const (
	str string = "book"
)

func main() {
	fmt.Println(halvesAreAlike(str))
}

func halvesAreAlike(s string) bool {
	vowelsMap := map[string]struct{}{"a": {}, "e": {}, "i": {}, "o": {}, "u": {}, "A": {}, "E": {}, "I": {}, "O": {}, "U": {}}

	firstCount, secondCount := 0, 0

	first, second := []rune(s[:len(s)/2]), []rune(s[len(s)/2:])

	for i := 0; i < len(s)/2; i++ {
		if _, ok := vowelsMap[string(first[i])]; ok {
			firstCount++
		}
		if _, ok := vowelsMap[string(second[i])]; ok {
			secondCount++
		}

	}
	return firstCount == secondCount
}
