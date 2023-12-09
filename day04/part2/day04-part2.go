package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type GameData struct {
	Wins map[int]int
	Have []int
}

var points = map[int]int{}

func (g GameData) GetPoints() float64 {
	var points float64 = 0

	for _, found := range g.Have {
		if _, ok := g.Wins[found]; ok {
			points += 1
		}
	}

	return points
}

var final int = 0

func GetAllPoints(points map[int]int) int {
	for _, point := range points {
		final += point
	}

	return final
}

func SumAll(nextIndex, n int) {
	for i := 0; i < n; i++ {
		if _, ok := points[nextIndex+i]; ok {
			points[nextIndex+i] += points[nextIndex-1]
		}
	}
}

func FillPoints(data []GameData) {
	for index := range data {
		points[index] = 1
	}
}

func DeepGame(filepath string) int {
	data := GetDataFromText(filepath)

	FillPoints(data)
	for index, game := range data {

		add := int(game.GetPoints())

		SumAll(index+1, add)
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

func main() {
	startTime := time.Now()
	result := DeepGame("puzzle_input")
	fmt.Println(result)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
