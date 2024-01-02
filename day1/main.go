package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		result, err := SumCalibrationNumber(text)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			fmt.Printf("Error tring to read line %s\n", text)
			break
		}

		total += result
	}

	fmt.Printf("total: %v\n", total)
}

func SumCalibrationNumber(text string) (int, error) {
	values := []rune{}
	i := 0

	for _, value := range text {
		if value >= '0' && value <= '9' {
			values = append(values, value)
			i++
		}
	}

	if i == 0 {
		return 0, errors.New("Invalid string")
	}
	str := string([]rune{values[0], values[i-1]})

	return strconv.Atoi(str)
}
