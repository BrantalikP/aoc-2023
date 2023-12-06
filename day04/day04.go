package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SumLines(n int) int {
	fmt.Println("SumLines: ", n)
	sum := 1
	// if n == 1 {
	// 	return 1
	// }

	for i := 0; i < n; i++ {
		fmt.Println("sum: ", sum)
		fmt.Println("i: ", i)
		sum += sum
	}
	return sum
}

type GameData struct {
	Wins map[int]int
	Have []int
}

var points = map[int]int{}

func DoubleNTimes(n float64) float64 {
	if n == 1 {
		return 1
	} else if n > 1 {
		return math.Pow(2, n-1)
	} else {
		return 0
	}
}

func (g GameData) GetPoints() float64 {
	var points float64 = 0

	for _, found := range g.Have {
		if _, ok := g.Wins[found]; ok {
			points += 1
		}
	}

	return points
}

func SumPoints(points []float64) float64 {
	var sum float64 = 0

	for _, point := range points {
		sum += point
	}

	return sum
}

var final int = 0

func GetAllPoints(points map[int]int) int {
	for _, point := range points {
		fmt.Println("point: ", point)
		final += point
	}

	return final
}

func SumAll(index, n int) {
	fmt.Println("SumAll: ", index, n)
	for i := 0; i < n; i++ {
		if _, ok := points[index+i]; ok {
			points[index+i] += points[index-1]
		} else {
			points[index+i] = 2
		}
	}
}

func DeepGame(filepath string) int {
	data := GetDataFromText(filepath)

	for index, game := range data {

		// fmt.Println(index)
		// fmt.Println(points)
		add := int(game.GetPoints())

		if add == 0 {
			final++
		} else {
			SumAll(index+1, add)
		}
	}

	return GetAllPoints(points)
}

func GetDataFromText(filepath string) []GameData {
	// Open file
	file, _ := os.Open(filepath + ".txt")
	scanner := bufio.NewScanner(file)

	var result []GameData

	// Scan file line by line
	for scanner.Scan() {
		lineData := strings.Split(scanner.Text(), "|")
		var data GameData
		data.Wins = make(map[int]int)

		// Get wins numbers
		wins := strings.Fields(lineData[0][8:])

		for _, win := range wins {
			winNum, _ := strconv.Atoi(win)
			data.Wins[winNum] = 1
		}

		// Get have numbers
		have := strings.Fields(lineData[1])
		data.Have = make([]int, len(have))

		for i, h := range have {
			haveNum, _ := strconv.Atoi(h)
			data.Have[i] = haveNum
		}

		result = append(result, data)
	}

	// Close file
	file.Close()

	return result
}

func GetSumPoints(filepath string) float64 {
	data := GetDataFromText(filepath)
	var points []float64

	for _, game := range data {
		points = append(points, game.GetPoints())
	}

	return SumPoints(points)
}

func main() {
	result := DeepGame("puzzle_input")
	fmt.Println(result)
}
