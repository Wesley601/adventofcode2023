package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Extractor interface {
	Extract(chunk string) (int, bool)
}

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("you pass -part1 or -part2"))
	}

	var solution Extractor

	switch os.Args[1] {
	case "-part1":
		solution = &FirstSolution{}
	case "-part2":
		solution = &SecondSolution{}
	default:
		panic(errors.New("invalid param"))
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := CalculateCalibrationFile(scanner, solution)

	fmt.Printf("result: %v\n", result)
}

func CalculateCalibrationFile(scanner *bufio.Scanner, extractor Extractor) int {
	var result int

	for scanner.Scan() {
		vv := 0
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if v, ok := extractor.Extract(line[i:]); ok {
				vv = v * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if v, ok := extractor.Extract(line[i:]); ok {
				vv += v
				break
			}
		}

		result += vv
	}

	return result
}
