package main

import (
	"aoc-2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputLines := util.GetInputLines("input.txt")
	fmt.Println("Part A: ", partA(inputLines))
	fmt.Println("Part B: ", partB(inputLines))
}

type cubeSet struct {
	red   int
	green int
	blue  int
}

type gameRecord struct {
	id      int
	samples []*cubeSet
}

func parseCubeSet(setStr string) *cubeSet {
	rReg := regexp.MustCompile(`(\d+) red`)
	gReg := regexp.MustCompile(`(\d+) green`)
	bReg := regexp.MustCompile(`(\d+) blue`)

	rCount, gCount, bCount := 0, 0, 0

	if rMatch := rReg.FindStringSubmatch(setStr); rMatch != nil {
		rCount, _ = strconv.Atoi(rMatch[1])
	}

	if gMatch := gReg.FindStringSubmatch(setStr); gMatch != nil {
		gCount, _ = strconv.Atoi(gMatch[1])
	}

	if bMatch := bReg.FindStringSubmatch(setStr); bMatch != nil {
		bCount, _ = strconv.Atoi(bMatch[1])
	}

	return &cubeSet{red: rCount, green: gCount, blue: bCount}
}

func parseGameRecord(gameStr string) *gameRecord {
	gameReg := regexp.MustCompile(`Game (\d+): (.+)`)
	gameMatch := gameReg.FindStringSubmatch(gameStr)
	gameId, _ := strconv.Atoi(gameMatch[1])

	var samples []*cubeSet
	sampleStrings := strings.Split(gameMatch[2], ";")

	for _, sampleString := range sampleStrings {
		samples = append(samples, parseCubeSet(sampleString))
	}

	return &gameRecord{id: gameId, samples: samples}
}

func (cSet cubeSet) possibleSampleFrom(other *cubeSet) bool {
	return other.red >= cSet.red && other.green >= cSet.green && other.blue >= cSet.blue
}

func (gRecord gameRecord) possibleRecordFrom(realSet *cubeSet) bool {
	for _, s := range gRecord.samples {
		if !s.possibleSampleFrom(realSet) {
			return false
		}
	}
	return true
}

func (g gameRecord) minSet() *cubeSet {
	rMin, gMin, bMin := 0, 0, 0

	for _, s := range g.samples {
		rMin = max(rMin, s.red)
		gMin = max(gMin, s.green)
		bMin = max(bMin, s.blue)
	}

	return &cubeSet{red: rMin, green: gMin, blue: bMin}
}

func partA(inputLines []string) string {
	realSet := &cubeSet{red: 12, green: 13, blue: 14}
	sum := 0
	for _, line := range inputLines {
		g := parseGameRecord(line)
		if g.possibleRecordFrom(realSet) {
			sum += g.id
		}
	}
	return fmt.Sprint(sum)
}
func partB(inputLines []string) string {
	sum := 0
	for _, line := range inputLines {
		minSet := parseGameRecord(line).minSet()
		sum += minSet.red * minSet.green * minSet.blue
	}
	return fmt.Sprint(sum)
}
