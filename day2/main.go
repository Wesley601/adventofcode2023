package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	solution := 0
	for scanner.Scan() {
		text := scanner.Text()
		solution += Part1(text)
	}

	fmt.Printf("solution: %v\n", solution)
}

func Part1(text string) int {
	games, gameID := SplitGame(text)

	gameIDInt, err := strconv.Atoi(gameID)
	if err != nil {
		panic(err)
	}

	for _, g := range games {
		if !IsPossibleGame(g) {
			return 0
		}
	}

	return gameIDInt
}

func Part2(text string) int {
	games, _ := SplitGame(text)

	max_red := 1
	max_green := 1
	max_blue := 1

	for _, vv := range games {
		for _, v := range strings.Split(vv, ",") {
			color, valueInt, err := GetColorAndValue(v)
			if err != nil {
				panic(err)
			}

			switch color {
			case "red":
				if valueInt > max_red {
					max_red = valueInt
				}
			case "green":
				if valueInt > max_green {
					max_green = valueInt
				}
			case "blue":
				if valueInt > max_blue {
					max_blue = valueInt
				}
			}
		}
	}

	return max_red * max_green * max_blue
}

func SplitGame(text string) ([]string, string) {
	games := strings.Split(text, ";")
	firstPos := games[0]

	tmp := strings.Split(firstPos, ":")

	gameID := strings.Split(tmp[0], " ")[1]
	games[0] = tmp[1]

	return games, gameID
}

func IsPossibleGame(s string) bool {
	for _, v := range strings.Split(s, ",") {
		color, valueInt, err := GetColorAndValue(v)
		if err != nil {
			panic(err)
		}

		if color == "red" && valueInt > MAX_RED {
			return false
		}

		if color == "green" && valueInt > MAX_GREEN {
			return false
		}
		if color == "blue" && valueInt > MAX_BLUE {
			return false
		}
	}

	return true
}

func GetColorAndValue(v string) (string, int, error) {
	tmp := strings.Split(strings.TrimPrefix(v, " "), " ")
	color := tmp[1]
	value := tmp[0]

	valueInt, err := strconv.Atoi(value)
	return color, valueInt, err
}
