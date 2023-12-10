package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var HandMapKeys = map[string]string{
	"high-card":       "0",
	"one-pair":        "1",
	"two-pairs":       "2",
	"three-of-a-kind": "3",
	"full-house":      "4",
	"four-of-a-kind":  "5",
	"five-of-a-kind":  "6",
}

var CardMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type CardHand struct {
	Hand  string
	Type  string
	Score int
}

func GetValue(cards string) int {
	sum := 0
	for i := 0; i < len(cards); i++ {
		parsedVal := string(cards[i])
		score := len(cards) - i
		sum += score * CardMap[parsedVal]
	}

	return sum
}

func SortCardHands(cards []CardHand) []CardHand {
	sort.Slice(cards, func(i, j int) bool {
		for k := 0; k < len(cards); k++ {
			scoreI := CardMap[string(cards[i].Hand[k])]
			scoreJ := CardMap[string(cards[j].Hand[k])]

			if scoreI != scoreJ {
				return scoreI < scoreJ
			}
		}

		return false
	})

	return cards
}

func SortAndMerge(cardsMap map[string][]CardHand) []CardHand {
	cardArray := []CardHand{}

	keys := make([]string, 0, len(cardsMap))
	for k := range cardsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sortedCards := SortCardHands(cardsMap[k])
		cardArray = append(cardArray, sortedCards...)
	}

	return cardArray
}

func SumWinner(cards []CardHand) int {
	sum := 0

	for i, cardHand := range cards {
		score := i + 1
		sum += score * cardHand.Score
	}

	return sum
}

func ScanAndFormatData(filepath string) (map[string][]CardHand, error) {
	handMap := map[string][]CardHand{}
	file, error := os.Open(filepath + ".txt")
	if error != nil {
		return nil, error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		stringArray := strings.Split(text, " ")

		handType := GetHand(stringArray[0])
		score, _ := strconv.Atoi(stringArray[1])

		newCardHand := CardHand{
			Hand:  stringArray[0],
			Type:  handType,
			Score: score,
		}

		key := HandMapKeys[handType]
		handMap[key] = append(handMap[key], newCardHand)
	}

	return handMap, nil
}

func GetHand(cards string) string {
	marked := map[string]int{}

	for _, val := range cards {
		parsedVal := string(val)
		if _, ok := marked[parsedVal]; ok {
			marked[parsedVal] += 1
		} else {
			marked[parsedVal] = 1
		}
	}

	switch len(marked) {
	case 1:
		return "five-of-a-kind"
	case 2:
		for _, val := range marked {
			if val == 4 {
				return "four-of-a-kind"
			}
		}
		return "full-house"
	case 3:
		for _, val := range marked {
			if val == 3 {
				return "three-of-a-kind"
			}
		}
		return "two-pairs"
	case 4:
		return "one-pair"
	case 5:
	default:
		return "high-card"
	}
	return ""
}

func main() {
	data, _ := ScanAndFormatData("puzzle_input")
	sortedData := SortAndMerge(data)
	result := SumWinner(sortedData)
	fmt.Println(result)
}
