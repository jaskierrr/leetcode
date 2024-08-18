package main

import "fmt"

const (
	str string = "textbook"
)

func main() {
	fmt.Println(halvesAreAlike(str))
}

func halvesAreAlike(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	
	vowelsMap := make(map[string]bool, len(vowels))
	for _, v := range vowels {
		vowelsMap[v] = true
	}

	firstCount := 0
	secondCount := 0

	first := make([]rune, 0, len(vowels)/2)
	second := make([]rune, 0, len(vowels)/2)

	first = append(first, []rune(s[:len(s)/2])...)
	second = append(second, []rune(s[len(s)/2:])...)

	for _, v := range first {
		if _, ok := vowelsMap[string(v)]; ok {
			firstCount++
		}
	}
	for _, v := range second {
		if _, ok := vowelsMap[string(v)]; ok {
			secondCount++
		}
	}

	fmt.Println(first, second)
	fmt.Println(firstCount, secondCount)

	return firstCount == secondCount
}
