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
	},
	"5": {
		{"T55J5", "four-of-a-kind", 684},
		{"KTJJT", "four-of-a-kind", 220},
		{"QQQJA", "four-of-a-kind", 483},
	},
}

var sorted_input = []CardHand{
	{"32T3K", "one-pair", 765},
	{"KK677", "two-pairs", 28},
	{"T55J5", "four-of-a-kind", 684},
	{"QQQJA", "four-of-a-kind", 483},
	{"KTJJT", "four-of-a-kind", 220},
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
	want := 5905

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestScanAndFormatData(t *testing.T) {
	got, _ := ScanAndFormatData("test_input")
	want := test_input

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
		{cards: "T55J5", want: "four-of-a-kind"},
		{cards: "KK677", want: "two-pairs"},
		{cards: "KTJJT", want: "four-of-a-kind"},
		{cards: "QQQJA", want: "four-of-a-kind"},
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
