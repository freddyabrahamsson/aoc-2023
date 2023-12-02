package main

import (
	"reflect"
	"testing"
)

var game1String string = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
var game1 gameRecord = gameRecord{
	id:      1,
	samples: []*cubeSet{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}},
}

var game2String string = "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
var game2 gameRecord = gameRecord{id: 2, samples: []*cubeSet{
	{blue: 1, green: 2},
	{green: 3, blue: 4, red: 1},
	{green: 1, blue: 1},
}}

var game3String string = "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
var game3 gameRecord = gameRecord{
	id:      3,
	samples: []*cubeSet{{green: 8, blue: 6, red: 20}, {blue: 5, red: 4, green: 13}, {green: 5, red: 1}},
}

func TestParseSampleA(t *testing.T) {
	input := "3 blue, 4 red"
	expected := &cubeSet{blue: 3, red: 4}
	actual := parseCubeSet(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}
func TestParseSampleB(t *testing.T) {
	input := "1 red, 2 green, 6 blue"
	expected := &cubeSet{red: 1, green: 2, blue: 6}
	actual := parseCubeSet(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}

func TestParseSampleC(t *testing.T) {
	input := "2 green"
	expected := &cubeSet{green: 2}
	actual := parseCubeSet(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}

func TestParseGameA(t *testing.T) {
	input := game1String
	expected := &game1
	actual := parseGameRecord(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}
func TestParseGameB(t *testing.T) {
	input := game2String
	expected := &game2
	actual := parseGameRecord(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}

func TestParseGameC(t *testing.T) {
	input := game3String
	expected := &game3
	actual := parseGameRecord(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}
func TestParseGameD(t *testing.T) {
	input := "Game 3322: 82 green, 64 blue, 20 red; 532 blue, 4 red, 131 green; 5 green, 1 red"
	expected := &gameRecord{
		id:      3322,
		samples: []*cubeSet{{green: 82, blue: 64, red: 20}, {blue: 532, red: 4, green: 131}, {green: 5, red: 1}},
	}
	actual := parseGameRecord(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}

func TestPossibleSampleFromA(t *testing.T) {
	sampleA := &cubeSet{blue: 3, red: 4}
	sampleB := &cubeSet{red: 12, green: 13, blue: 14}
	expected := true
	actual := sampleA.possibleSampleFrom(sampleB)
	if expected != actual {
		t.Errorf("%#v should be a possible sample from %#v", sampleA, sampleB)
	}
}

func TestPossibleSampleFromB(t *testing.T) {
	sampleA := &cubeSet{blue: 3, red: 4, green: 14}
	sampleB := &cubeSet{red: 12, green: 13, blue: 14}
	expected := false
	actual := sampleA.possibleSampleFrom(sampleB)
	if expected != actual {
		t.Errorf("%#v should not be a possible sample from %#v", sampleA, sampleB)
	}
}

func TestPossibleRecordFromA(t *testing.T) {
	realSet := cubeSet{red: 12, green: 13, blue: 14}
	recorded := game1
	expected := true
	actual := recorded.possibleRecordFrom(&realSet)
	if expected != actual {
		t.Errorf("%#v should be a possible game from %#v", recorded, realSet)
	}
}

func TestPossibleRecordFromB(t *testing.T) {
	realSet := cubeSet{red: 12, green: 13, blue: 14}
	recorded := game3
	expected := false
	actual := recorded.possibleRecordFrom(&realSet)
	if expected != actual {
		t.Errorf("%#v should not be a possible game from %#v", recorded, realSet)
	}
}

func TestMinSetA(t *testing.T) {
	recorded := game1
	expected := &cubeSet{red: 4, green: 2, blue: 6}
	actual := recorded.minSet()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v\nas minimal set for %#v\n got %#v", expected, recorded, *actual)
	}
}

func TestMinSetB(t *testing.T) {
	recorded := game2
	expected := &cubeSet{red: 1, green: 3, blue: 4}
	actual := recorded.minSet()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v\nas minimal set for %#v\n got %#v", expected, recorded, *actual)
	}
}

func TestMinSetC(t *testing.T) {
	recorded := game3
	expected := &cubeSet{red: 20, green: 13, blue: 6}
	actual := recorded.minSet()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v\nas minimal set for %#v\n got %#v", expected, recorded, *actual)
	}
}

func TestPartA(t *testing.T) {
	inputLines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := "8"
	actual := partA(inputLines)

	if actual != expected {
		t.Errorf("Part A failed, expected '%s', got '%s'", expected, actual)
	}
}
func TestPartB(t *testing.T) {
	inputLines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := "2286"
	actual := partB(inputLines)

	if actual != expected {
		t.Errorf("Part B failed, expected '%s', got '%s'", expected, actual)
	}
}
