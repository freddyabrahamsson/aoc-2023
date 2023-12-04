package main

import (
	"aoc-2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputLines := util.GetInputLines("input.txt")
	fmt.Println("Part A: ", partA(inputLines))
	fmt.Println("Part B: ", partB(inputLines))
}

type card struct {
	cardID         int
	winningNumbers map[int]struct{}
	myNumbers      map[int]struct{}
}

func parseSpaceSeparatedNumbers(numString string) *map[int]struct{} {
	nums := make(map[int]struct{})
	for _, winningNumStr := range strings.Split(numString, " ") {
		num, err := strconv.Atoi(winningNumStr)
		if err == nil {
			nums[num] = struct{}{}
		}
	}
	return &nums
}

func parseCard(cardStr string) card {
	cardReg := regexp.MustCompile(`Card\s+(\d+):(.+)\|(.+)`)
	card := card{}
	if cardMatch := cardReg.FindStringSubmatch(cardStr); cardMatch != nil {
		id, _ := strconv.Atoi(cardMatch[1])
		card.cardID = id
		card.winningNumbers = *parseSpaceSeparatedNumbers(cardMatch[2])
		card.myNumbers = *parseSpaceSeparatedNumbers(cardMatch[3])
	}
	return card
}

func (c card) points() int {
	winners := c.nWinners()
	if winners == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(winners-1)))
	}
}

func (c card) nWinners() int {
	winners := 0
	for k := range c.myNumbers {
		if _, inMap := c.winningNumbers[k]; inMap {
			winners += 1
		}
	}
	return winners
}

func partA(inputLines []string) string {
	sum := 0
	for _, line := range inputLines {
		card := parseCard(line)
		sum += card.points()
	}
	return fmt.Sprint(sum)
}
func partB(inputLines []string) string {
	cardCount := map[int]int{}
	cards := make([]card, len(inputLines))
	for n, line := range inputLines {
		card := parseCard(line)
		cardCount[card.cardID] = 1
		cards[n] = card
	}
	for _, card := range cards {
		nWinners := card.nWinners()
		nCopies := cardCount[card.cardID]
		for i := 1; i <= nWinners; i++ {
			cardCount[card.cardID+i] += nCopies
		}
	}
	sum := 0
	for _, v := range cardCount {
		sum += v
	}
	return fmt.Sprint(sum)
}
