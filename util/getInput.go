package util

import (
	"bufio"
	"fmt"
	"os"
)

// Read an input file into a slice of strings, with each line in the input file as an element in the slice.
func GetInputLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return lines
}
