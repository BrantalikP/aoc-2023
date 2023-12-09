package main

import (
	"reflect"
	"testing"
)

var TestData = []Race{
	{Time: 7, Distance: 9}, {Time: 15, Distance: 40}, {Time: 30, Distance: 200},
}

const (
	timestring     = "     7  15   30"
	distancestring = "  9  40  200"
)

func TestScanData(t *testing.T) {
	got, _ := ScanData("test_input")
	want := TestData

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRaceData(t *testing.T) {
	t.Run("timestring", func(t *testing.T) {
		got := GetRaceData(timestring)
		want := []string{"7", "15", "30"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("distancestring", func(t *testing.T) {
		got := GetRaceData(distancestring)
		want := []string{"9", "40", "200"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAndCal(t *testing.T) {
	got := SumAndCalData(TestData)
	want := 288

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
