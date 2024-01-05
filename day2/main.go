package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var gameID string
		text := scanner.Text()

		values := strings.Split(text, ":")

		gameID = strings.Split(values[0], " ")[1]

		strings.Split(values[1], ", ")
	}
}
