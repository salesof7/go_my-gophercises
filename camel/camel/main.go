package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	answerStr := 1

	resolveByRune(input, &answer)
	resolveByString(input, &answerStr)

	fmt.Printf("%d\n%d", answer, answerStr)
}

func resolveByRune(input string, c *int) {
	for _, r := range input {
		min, max := 'A', 'Z'
		if r >= min && r <= max {
			*c++
		}
	}
}

func resolveByString(input string, c *int) {
	for _, r := range input {
		str := string(r)
		if strings.ToUpper(str) == str {
			*c++
		}
	}
}
