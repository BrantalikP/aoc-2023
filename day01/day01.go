package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCalibrationValue(input string) int {
	var sum []string
	// LOL ^2
	replaced := ReplaceWrittenNumbers(input)
	parsedInput := ReplaceWrittenNumbers(replaced)
	for _, char := range parsedInput {
		if char > 48 && char < 58 {

			number := string(char)
			sum = append(sum, number)
		}
	}

	value := ""

	length := len(sum)
	switch length {
	case 0:
		value = "0"
	case 1:
		value = sum[0] + sum[0]
	default:
		value = sum[0] + sum[len(sum)-1]

	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return parsedValue
}

func GetSumOfCalibrationValues(input []string) int {
	sum := 0
	for _, string := range input {
		value := GetCalibrationValue(string)
		sum += value
	}

	return sum
}

func ScanAndReturn(filepath string) []string {
	var lines []string
	readFile, err := os.Open(filepath + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	readFile.Close()
	return lines
}

// LOL
func ReplaceWrittenNumbers(line string) string {
	replacer := strings.NewReplacer(
		"one",
		"o1e",
		"two",
		"t2o",
		"three",
		"t3e",
		"four",
		"f4r",
		"five",
		"f5e",
		"six",
		"s6x",
		"seven",
		"s7n",
		"eight",
		"e8t",
		"nine",
		"n9e",
		"zero",
		"z0o",
	)

	return replacer.Replace(line)
}

func GetResult(filepathname string) int {
	lines := ScanAndReturn(filepathname)
	result := GetSumOfCalibrationValues(lines)
	return result
}

func main() {
	result1 := GetResult("puzzleInput")
	fmt.Println(result1)
}
