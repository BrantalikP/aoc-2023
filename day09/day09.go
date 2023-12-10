package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanData(filename string) ([][]int, error) {
	data := [][]int{}
	file, error := os.Open(filename + ".txt")
	if error != nil {
		return nil, error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		stringArray := strings.Split(text, " ")

		var numberArray []int
		for _, str := range stringArray {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
				return nil, error
			}
			numberArray = append(numberArray, number)
		}

		data = append(data, numberArray)
	}

	return data, nil
}

func GetMap(nums []int) [][]int {
	dataMap := [][]int{nums}

	allZero := false

	for !allZero {

		temSum := []int{}
		length := len(dataMap[len(dataMap)-1]) - 1
		allZero = true
		for j := 0; j < length; j++ {

			number1 := dataMap[len(dataMap)-1][j]
			number2 := dataMap[len(dataMap)-1][j+1]

			result := number2 - number1
			temSum = append(temSum, result)

			if result != 0 {
				allZero = false
			}

		}
		if !allZero {
			dataMap = append(dataMap, temSum)
		}
	}

	return dataMap
}

func CalculateHistory(nums [][]int) int {
	sum := 0
	for _, val := range nums {

		number := val[len(val)-1]
		sum += number
	}
	return sum
}

func CalculateReversedHistory(nums [][]int) int {
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {

		number := nums[i][0]
		newNumber := number - sum
		sum = newNumber

	}
	return sum
}

func SumAll(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	return sum
}

func main() {
	historyValues := []int{}
	data, _ := ScanData("test_input")

	for _, nums := range data {
		dataMap := GetMap(nums)
		historyVal := CalculateReversedHistory(dataMap)

		historyValues = append(historyValues, historyVal)
	}

	result := SumAll(historyValues)
	fmt.Println(result)
}
