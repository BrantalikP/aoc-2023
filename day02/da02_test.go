package main

import (
	"testing"
)

var mapOfGames = GameData{
	"1": {
		{"blue": 3, "red": 4},
		{"blue": 6, "green": 2, "red": 1},
		{"green": 2},
	},
	"2": {
		{"blue": 1, "green": 2},
		{"blue": 4, "green": 3, "red": 1},
		{"blue": 1, "green": 1},
	},
	"3": {
		{"blue": 6, "green": 8, "red": 20},
		{"blue": 5, "green": 13, "red": 4},
		{"green": 5, "red": 1},
	},
	"4": {
		{"blue": 6, "green": 1, "red": 3},
		{"green": 3, "red": 6},
		{"blue": 15, "green": 3, "red": 14},
	},
	"5": {
		{"blue": 1, "green": 3, "red": 6},
		{"blue": 2, "green": 2, "red": 1},
	},
}

//	func TestParsing(t *testing.T) {
//		t.Run("TestParsing", func(t *testing.T) {
//			// ParseData("testInput")
//		})
//	}
//
//	func TestSearch(t *testing.T) {
//		t.Run("TestSearch", func(t *testing.T) {
//			Search(mapOfGames)
//		})
//	}
func TestSearchAndMultiply(t *testing.T) {
	t.Run("TestSearch", func(t *testing.T) {
		SearchAndMultiply(mapOfGames)
	})
}
