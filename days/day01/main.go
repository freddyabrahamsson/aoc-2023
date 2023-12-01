package main

import (
	"aoc-2023/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	inputLines := util.GetInputLines("input.txt")
	fmt.Println("Part A: ", partA(inputLines))
	fmt.Println("Part B: ", partB(inputLines))
}

func partA(inputLines []string) string {
	sum := 0
	for _, inputLine := range inputLines {
		digits := extractDigits(inputLine)
		digitA := digits[0]
		digitB := digits[len(digits)-1]
		sum += 10*digitA + digitB
	}
	return fmt.Sprint(sum)
}
func partB(inputLines []string) string {
	sum := 0
	for _, inputLine := range inputLines {
		inputLine = substituteDigits(inputLine)
		digits := extractDigits(inputLine)
		digitA := digits[0]
		digitB := digits[len(digits)-1]
		sum += 10*digitA + digitB
	}
	return fmt.Sprint(sum)
}

// Returns a slice containing all digits found in a given string.
func extractDigits(inputLine string) []int {
	var digits []int
	for _, ch := range inputLine {
		digit, err := strconv.Atoi(string(ch))
		if err == nil {
			digits = append(digits, digit)
		}
	}
	return digits
}

// Inserts a digit representation
func substituteDigits(text string) string {
	var digitStrings []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for idx, digitString := range digitStrings {
		r := regexp.MustCompile(digitString)
		text = r.ReplaceAllString(text, fmt.Sprint(digitString, idx+1, digitString))
	}
	return text
}
