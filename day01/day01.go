package main

import (
	"strconv"
)

func GetCalibrationValue(input string) int {
	var sum []string
	for _, char := range input {
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
	return []string{"", ""}
}

func main() {
	GetCalibrationValue("1abc2")
}
