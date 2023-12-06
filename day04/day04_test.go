package main

import (
	"reflect"
	"testing"
)

var gameData = []GameData{
	{
		Wins: map[int]int{41: 1, 48: 1, 83: 1, 86: 1, 17: 1},
		Have: []int{83, 86, 6, 31, 17, 9, 48, 53},
	},
	{
		Wins: map[int]int{13: 1, 32: 1, 20: 1, 16: 1, 61: 1},
		Have: []int{61, 30, 68, 82, 17, 32, 24, 19},
	},
	{
		Wins: map[int]int{1: 1, 21: 1, 53: 1, 59: 1, 44: 1},
		Have: []int{69, 82, 63, 72, 16, 21, 14, 1},
	},
	{
		Wins: map[int]int{41: 1, 92: 1, 73: 1, 84: 1, 69: 1},
		Have: []int{59, 84, 76, 51, 58, 5, 54, 83},
	},
	{
		Wins: map[int]int{87: 1, 83: 1, 26: 1, 28: 1, 32: 1},
		Have: []int{88, 30, 70, 12, 93, 22, 82, 36},
	},
	{
		Wins: map[int]int{31: 1, 18: 1, 13: 1, 56: 1, 72: 1},
		Have: []int{74, 77, 10, 23, 35, 67, 36, 11},
	},
}

// func TestSumLines(t *testing.T) {
// 	got := SumLines(2)
// 	var want int = 2
//
// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

func TestGetDataFromText(t *testing.T) {
	got := GetDataFromText("test_data")
	want := gameData

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetPoints(t *testing.T) {
	got := gameData[0].GetPoints()
	var want float64 = 8

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDoubleNTimes(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := DoubleNTimes(1)

		var want float64 = 1

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test 4", func(t *testing.T) {
		got := DoubleNTimes(4)

		var want float64 = 8

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test 0", func(t *testing.T) {
		got := DoubleNTimes(0)

		var want float64 = 0

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetSumPoints(t *testing.T) {
	got := GetSumPoints("test_data")
	var want float64 = 13

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
