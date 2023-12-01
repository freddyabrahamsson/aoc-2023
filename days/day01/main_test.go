package main

import (
	"reflect"
	"testing"
)

func TestExtractDigitsA(t *testing.T) {
	text := "1abc2"
	digits := []int{1, 2}
	if !reflect.DeepEqual(extractDigits(text), digits) {
		t.Fatalf("%s should contain digits %#v", text, digits)
	}
}
func TestExtractDigitsB(t *testing.T) {
	text := "pqr3stu8vwx"
	digits := []int{3, 8}
	if !reflect.DeepEqual(extractDigits(text), digits) {
		t.Fatalf("%s should contain digits %#v", text, digits)
	}
}

func TestExtractDigitsC(t *testing.T) {
	text := "a1b2c3d4e5f"
	digits := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(extractDigits(text), digits) {
		t.Fatalf("%s should contain digits %#v", text, digits)
	}
}
func TestExtractDigitsD(t *testing.T) {
	text := "treb7uchet"
	digits := []int{7}
	if !reflect.DeepEqual(extractDigits(text), digits) {
		t.Fatalf("%s should contain digits %#v", text, digits)
	}
}

func TestPartA(t *testing.T) {
	inputLines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	expected := "142"
	actual := partA(inputLines)

	if actual != expected {
		t.Errorf("Part A failed, expected '%s', got '%s'", expected, actual)
	}
}
func TestPartB(t *testing.T) {
	inputLines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := "281"
	actual := partB(inputLines)

	if actual != expected {
		t.Errorf("Part B failed, expected '%s', got '%s'", expected, actual)
	}
}
