package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

const (
	timeSeparator     = "Time:"
	distanceSeparator = "Distance:"
)

func ScanData(filename string) ([]Race, error) {
	file, error := os.Open(filename + ".txt")
	if error != nil {
		return nil, errors.New("Fuck")
	}

	return GetData(file)
}

func GetData(body io.Reader) ([]Race, error) {
	scanner := bufio.NewScanner(body)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	timeValues := GetRaceData(readMetaLine(timeSeparator))
	distanceValues := GetRaceData(readMetaLine(distanceSeparator))

	data := ConcatData(timeValues, distanceValues)

	return data, nil
}

func ConcatData(timeValues, distanceValues []string) []Race {
	data := []Race{}
	for index, val := range timeValues {
		timeValue, _ := strconv.Atoi(val)
		distanceValue, _ := strconv.Atoi(distanceValues[index])

		data = append(data, Race{Time: timeValue, Distance: distanceValue})

	}

	return data
}

func GetRaceData(text string) []string {
	re := regexp.MustCompile("[0-9]+")

	timeValues := re.FindAllString(text, -1)

	return timeValues
}

func SumAndCalData(raceData []Race) int {
	sum := 1
	for _, data := range raceData {
		value := CalRecords(data)
		sum *= value
	}
	return sum
}

func CalRecords(data Race) int {
	won := 0

	for i := 0; i < data.Time; i++ {
		speed := i
		remaingTime := data.Time - i

		distance := speed * remaingTime

		if distance > data.Distance {
			won++
		}
	}

	return won
}

func main() {
	recordData, _ := ScanData("puzzle_input")
	result := SumAndCalData(recordData)

	fmt.Println(result)
}
