package main

import (
	"fmt"
	"reflect"
	"testing"
)

var test_input = [][]int{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
}

var test_input_string = [][]string{
	{"0", "3", "6", "9", "12", "15"},
	{"1", "3", "6", "10", "15", "21"},
	{"10", "13", "16", "21", "30", "45"},
}

var firstLine = [][]int{
	{0, 3, 6, 9, 12, 15},
	{3, 3, 3, 3, 3},
}

var secondLine = [][]int{
	{1, 3, 6, 10, 15, 21},
	{2, 3, 4, 5, 6},
	{1, 1, 1, 1},
}

var thirdLine = [][]int{
	{10, 13, 16, 21, 30, 45},
	{3, 3, 5, 9, 15},
	{0, 2, 4, 6},
	{2, 2, 2},
}

func TestScanData(t *testing.T) {
	got, _ := ScanData("test_input")
	want := test_input

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetMap(t *testing.T) {
	areaTests := []struct {
		array  []int
		result [][]int
	}{
		{array: test_input[0], result: firstLine},
		{array: test_input[1], result: secondLine},
		{array: test_input[2], result: thirdLine},
	}

	for _, tt := range areaTests {
		t.Run(fmt.Sprintf("%d", tt.result), func(t *testing.T) {
			got := GetMap(tt.array)

			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("got %v, want %v", got, tt.result)
			}
		})
	}
}

func TestCalculateHistory(t *testing.T) {
	areaTests := []struct {
		array  [][]int
		result int
	}{
		{array: firstLine, result: 18},
		{array: secondLine, result: 28},
		{array: thirdLine, result: 68},
	}

	for _, tt := range areaTests {
		t.Run(fmt.Sprintf("%d", tt.result), func(t *testing.T) {
			got := CalculateHistory(tt.array)

			if got != tt.result {
				t.Errorf("got %d, want %d", got, tt.result)
			}
		})
	}
}

func TestCalculateReversedHistory(t *testing.T) {
	areaTests := []struct {
		array  [][]int
		result int
	}{
		{array: firstLine, result: -3},
		{array: secondLine, result: 0},
		{array: thirdLine, result: 5},
	}

	for _, tt := range areaTests {
		t.Run(fmt.Sprintf("%d", tt.result), func(t *testing.T) {
			got := CalculateReversedHistory(tt.array)

			if got != tt.result {
				t.Errorf("got %d, want %d", got, tt.result)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{18, 28, 68})
	want := 114

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
