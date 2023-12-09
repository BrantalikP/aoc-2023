package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Num struct {
	Value int
	Coor  []int
	Line  int
}

func Map(filepath string) (num []Num, chars [][]bool) {
	file, _ := os.Open(filepath + ".txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []Num
	var charMap [][]bool

	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		charMap = append(charMap, make([]bool, len(text)))

		re := regexp.MustCompile("[0-9]+")
		charRe := regexp.MustCompile("[^0-9.]")
		numArr := re.FindAllString(text, -1)
		numIndices := re.FindAllStringIndex(text, -1)

		charArr := charRe.FindAllString(text, -1)
		charIndices := charRe.FindAllStringIndex(text, -1)

		for i := range charArr {
			charMap[line][charIndices[i][0]] = true
		}

		for i, v := range numArr {
			parsedValue, _ := strconv.Atoi(v)
			number := Num{Value: parsedValue, Coor: numIndices[i], Line: line}
			numbers = append(numbers, number)

		}
		line++
	}

	return numbers, charMap
}

func FindOverlap(nums []Num, chars [][]bool) int {
	sum := 0
	for _, val := range nums {
		for i := val.Coor[0] - 1; i < val.Coor[1]+1; i++ {
			if i >= 0 && i+1 <= len(chars[val.Line]) {
				char1 := false
				char2 := chars[val.Line][i]
				char3 := false

				if val.Line-1 >= 0 {
					char1 = chars[val.Line-1][i]
				}
				if val.Line+1 < len(chars) {
					char3 = chars[val.Line+1][i]
				}
				if char1 == true || char2 == true || char3 == true {
					sum += val.Value
					break
				}
			}
		}
	}
	return sum
}

func main() {
	numbers, chars := Map("puzzle_input")
	result := FindOverlap(numbers, chars)
	fmt.Println(result)
}
