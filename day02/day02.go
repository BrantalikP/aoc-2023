package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	blue  = 14
	red   = 12
	green = 13
)

type (
	GameCounts map[string]int
	GameData   map[string][]GameCounts
)

func ParseData(pathname string) GameData {
	file, _ := os.Open(pathname + ".txt")
	scanner := bufio.NewScanner(file)

	gamesData := make(GameData)

	for scanner.Scan() {
		gameData := strings.Split(scanner.Text(), ":")
		gameId := gameData[0][5:]

		rounds := strings.Split(gameData[1], ";")
		var gameCountsSlice []GameCounts

		for _, round := range rounds {
			round = strings.TrimSpace(round)
			gameCounts := make(GameCounts)

			elements := strings.Split(round, ",")
			for _, element := range elements {
				element = strings.TrimSpace(element)
				color := strings.Split(element, " ")[1]
				count, _ := strconv.Atoi(strings.Split(element, " ")[0])

				gameCounts[color] = count
			}

			gameCountsSlice = append(gameCountsSlice, gameCounts)
		}

		gamesData[gameId] = gameCountsSlice
	}

	file.Close()
	return gamesData
}

func Search(gamesData GameData) {
	colors := map[string]int{"blue": blue, "red": red, "green": green}

	sum := 0
	for gameID, game := range gamesData {
		fmt.Printf("Game ID: %s\n", gameID)

		allRoundsInLimit := true // Flag to check if all rounds pass
		for round, counts := range game {
			fmt.Printf("\tRound %d: \n", round+1)

			for color, count := range counts {
				// If count is greater than maximum allowed for current color, set flag false and move to next game
				if maxAllowed, ok := colors[color]; ok && count > maxAllowed {
					allRoundsInLimit = false
					break
				}

				// Right align for better formatting
				fmt.Printf("\t\t%-5s: %d \n", strings.Title(color), count)
			}
			if !allRoundsInLimit {
				break
			}
		}

		// If game passed, add GameID to sum
		if allRoundsInLimit {
			gameIDInt, _ := strconv.Atoi(gameID)
			sum += gameIDInt
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func SearchAndMultiply(gamesData GameData) {
	sum := 0

	for gameID, game := range gamesData {
		fmt.Printf("Game ID: %s\n", gameID)
		maxValues := map[string]int{"blue": 1, "red": 1, "green": 1}
		subSum := 1

		for _, counts := range game {
			// fmt.Printf("\tRound %d: \n", round+1)

			for color, count := range counts {
				// If count is greater than maximum allowed for current color, set flag false and move to next game
				previousValue := maxValues[color]
				if count > previousValue {
					maxValues[color] = count
				}

				// Right align for better formatting
				// fmt.Printf("\t\t%-5s: %d \n", strings.Title(color), count)
			}
		}
		fmt.Println("one game")
		fmt.Println(maxValues)
		for color, count := range maxValues {
			fmt.Println(color, count)
			subSum *= count
		}

		sum += subSum
	}
	// fmt.Println(maxValues)
	// for color, count := range maxValues {
	// 	fmt.Println(color, count)
	// }
	fmt.Printf("Sum: %d\n", sum)
}

func main() {
	gameData := ParseData("puzzleInput")
	SearchAndMultiply(gameData)
}
