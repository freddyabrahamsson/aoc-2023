package util

import (
	"reflect"
	"testing"
)

func TestGetInputLines(t *testing.T) {
	expected := []string{"This is line one", "This is line two"}
	actual := GetInputLines("sample.txt")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%#v', got '%#v'", expected, actual)
	}
}
