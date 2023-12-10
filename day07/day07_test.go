package main

import (
	"reflect"
	"testing"
)

// func TestScanData (t *testing.T) {
// 	got := ScanData("test_input")
// 	want := {}
//
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

var test_input = map[string][]CardHand{
	"1": {
		{"32T3K", "one-pair", 765},
	},
	"2": {
		{"KK677", "two-pairs", 28},
		{"KTJJT", "two-pairs", 220},
	},
	"3": {
		{"T55J5", "three-of-a-kind", 684},
		{"QQQJA", "three-of-a-kind", 483},
	},
}

var test_input2 = map[string][]CardHand{
	"1": {
		{"32T3K", "one-pair", 765},
	},
	"2": {
		{"KK677", "two-pairs", 28},
		{"KTJJT", "two-pairs", 220},
	},
	"3": {
		{"T55J5", "three-of-a-kind", 684},
		{"QQQJA", "three-of-a-kind", 483},
	},
}

var sorted_input = []CardHand{
	{"32T3K", "one-pair", 765},
	{"KTJJT", "two-pairs", 220},
	{"KK677", "two-pairs", 28},
	{"T55J5", "three-of-a-kind", 684},
	{"QQQJA", "three-of-a-kind", 483},
}

func TestSortAndMerge(t *testing.T) {
	got := SortAndMerge(test_input)
	want := sorted_input

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v got, %v want", got, want)
	}
}

func TestSumWinner(t *testing.T) {
	got := SumWinner(sorted_input)
	want := 6440

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestScanAndFormatData(t *testing.T) {
	got, _ := ScanAndFormatData("test_input")
	want := test_input2

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v got, want %v", got, want)
	}
}

func TestGetHand(t *testing.T) {
	areaTests := []struct {
		cards string
		want  string
	}{
		{cards: "32T3K", want: "one-pair"},
		{cards: "T55J5", want: "three-of-a-kind"},
		{cards: "KK677", want: "two-pairs"},
		{cards: "KTJJT", want: "two-pairs"},
		{cards: "QQQJA", want: "three-of-a-kind"},
	}

	for _, tt := range areaTests {
		t.Run(tt.want, func(t *testing.T) {
			got := GetHand(tt.cards)

			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

// func TestGetValue(t *testing.T) {
// 	areaTests := []struct {
// 		cards string
// 		want  int
// 	}{
// 		{cards: "32T3K", want: 72},
// 		{cards: "T55J5", want: 112},
// 		// {cards: "KK677", want: "two-pairs"},
// 		// {cards: "KTJJT", want: "two-pairs"},
// 		// {cards: "QQQJA", want: "three-of-a-kind"},
// 	}
//
// 	for _, tt := range areaTests {
// 		t.Run(tt.cards, func(t *testing.T) {
// 			got := GetValue(tt.cards)
//
// 			if got != tt.want {
// 				t.Errorf("got %d, want %d", got, tt.want)
// 			}
// 		})
// 	}
// }
