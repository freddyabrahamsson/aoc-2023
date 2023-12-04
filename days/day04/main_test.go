package main

import (
	"reflect"
	"testing"
)

var card1Str string = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
var card1 card = card{
	cardID: 1,
	winningNumbers: map[int]struct{}{
		41: {}, 48: {}, 83: {}, 86: {}, 17: {},
	},
	myNumbers: map[int]struct{}{
		83: {}, 86: {}, 6: {}, 31: {}, 17: {}, 9: {}, 48: {}, 53: {},
	},
}

var card2Str string = "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
var card2 card = card{
	cardID: 2,
	winningNumbers: map[int]struct{}{
		13: {}, 32: {}, 20: {}, 16: {}, 61: {},
	},
	myNumbers: map[int]struct{}{
		61: {}, 30: {}, 68: {}, 82: {}, 17: {}, 32: {}, 24: {}, 19: {},
	},
}

func TestParseCardA(t *testing.T) {
	input := card1Str
	expected := card1
	actual := parseCard(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}
func TestParseCardB(t *testing.T) {
	input := card2Str
	expected := card2
	actual := parseCard(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect parse for '%s', expected %#v, got %#v", input, expected, actual)
	}
}
func TestCardPointsA(t *testing.T) {
	input := card1
	expected := 8
	actual := input.points()
	if expected != actual {
		t.Errorf("Expected '%#v' to give '%d' points, got '%d'", input, expected, actual)
	}
}

func TestPartA(t *testing.T) {
	inputLines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	expected := "13"
	actual := partA(inputLines)

	if actual != expected {
		t.Errorf("Part A failed, expected '%s', got '%s'", expected, actual)
	}
}
func TestPartB(t *testing.T) {
	inputLines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	expected := "30"
	actual := partB(inputLines)

	if actual != expected {
		t.Errorf("Part B failed, expected '%s', got '%s'", expected, actual)
	}
}
