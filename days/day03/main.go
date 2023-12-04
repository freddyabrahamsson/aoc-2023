package main

import (
	"aoc-2023/util"
	"fmt"
	"strconv"
)

func main() {
	inputLines := util.GetInputLines("input.txt")
	fmt.Println("Part A: ", partA(inputLines))
	fmt.Println("Part B: ", partB(inputLines))
}

type number struct {
	value           int
	leftCoordinate  [2]int
	rightCoordinate [2]int
}

type grid struct {
	content [][]rune
	numbers []number
}

func parseGridAndNumbers(inputLines []string) grid {
	nLines := len(inputLines)
	nCols := len(inputLines[0])
	g := grid{content: make([][]rune, nLines), numbers: []number{}}

	for lineNumber, lineContent := range inputLines {
		g.content[lineNumber] = make([]rune, nCols)

		readingNumber := false
		var currentNumber number

		for charPos, char := range lineContent {
			num, err := strconv.Atoi(string(char))
			if err == nil { // Character was a digit.
				if readingNumber {
					currentNumber.value = 10*currentNumber.value + num
				} else { // Found new number
					currentNumber = number{leftCoordinate: [2]int{charPos, lineNumber}, value: num}
					readingNumber = true
				}
			} else { // Character was not a digit.
				if readingNumber {
					currentNumber.rightCoordinate = [2]int{charPos - 1, lineNumber}
					g.numbers = append(g.numbers, currentNumber)
					readingNumber = false
				}
			}
			g.content[lineNumber][charPos] = char
		}
		if readingNumber {
			currentNumber.rightCoordinate = [2]int{len(lineContent) - 1, lineNumber}
			g.numbers = append(g.numbers, currentNumber)
		}
	}
	return g
}

func (g grid) getNumNeighbours(leftCoordinate [2]int, rightCoordinate [2]int) []rune {
	neighbours := []rune{}
	for _, coord := range g.getNumNeighbourCoords(leftCoordinate, rightCoordinate) {
		neighbours = append(neighbours, g.content[coord[1]][coord[0]])
	}
	return neighbours
}

func (g grid) getNumNeighbourCoords(leftCoord [2]int, rightCoord [2]int) [][2]int {
	coords := [][2]int{}
	y := leftCoord[1]
	xLeft := leftCoord[0]
	xRight := rightCoord[0]

	hasLeft := xLeft > 0
	hasRight := xRight < len(g.content[0])-1
	hasAbove := y > 0
	hasBelow := y < len(g.content)-1

	if hasAbove {
		for x := xLeft; x <= xRight; x++ {
			coords = append(coords, [2]int{x, y - 1})
		}
		if hasLeft {
			coords = append(coords, [2]int{xLeft - 1, y - 1})
		}
		if hasRight {
			coords = append(coords, [2]int{xRight + 1, y - 1})

		}
	}
	if hasBelow {
		for x := xLeft; x <= xRight; x++ {
			coords = append(coords, [2]int{x, y + 1})
		}
		if hasLeft {
			coords = append(coords, [2]int{xLeft - 1, y + 1})
		}
		if hasRight {
			coords = append(coords, [2]int{xRight + 1, y + 1})
		}
	}
	if hasLeft {
		coords = append(coords, [2]int{xLeft - 1, y})
	}
	if hasRight {
		coords = append(coords, [2]int{xRight + 1, y})
	}

	return coords
}

func isPartSymbol(r rune) bool {
	_, failedToParseNumber := strconv.Atoi(string(r))
	isDot := r == '.'
	return !isDot && failedToParseNumber != nil
}

func partA(inputLines []string) string {
	g := parseGridAndNumbers(inputLines)
	sum := 0
	for _, num := range g.numbers {
		neighbours := g.getNumNeighbours(num.leftCoordinate, num.rightCoordinate)
		for _, neighbour := range neighbours {
			if isPartSymbol(neighbour) {
				sum += num.value
				break
			}
		}
	}
	return fmt.Sprint(sum)
}

func partB(inputLines []string) string {
	g := parseGridAndNumbers(inputLines)
	var asterisks map[int]map[int][]int = make(map[int]map[int][]int)

	for _, num := range g.numbers {
		for _, coord := range g.getNumNeighbourCoords(num.leftCoordinate, num.rightCoordinate) {
			x, y := coord[0], coord[1]

			if g.content[y][x] == '*' {
				if _, exists := asterisks[y]; !exists {
					asterisks[y] = make(map[int][]int)
				}
				asterisks[y][x] = append(asterisks[y][x], num.value)
			}
		}
	}

	sum := 0
	for _, row := range asterisks {
		for _, list := range row {
			if len(list) == 2 {
				sum += list[0] * list[1]
			}
		}
	}

	return fmt.Sprint(sum)
}
