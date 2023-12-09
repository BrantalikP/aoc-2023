package main

import (
	"reflect"
	"testing"
)

var test_map = [][]string{
	{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
	{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
	{".", ".", "3", "5", ".", ".", "6", "3", "3", "."},
}

var testNum = []Num{
	{Value: 467, Coor: []int{0, 3}, Line: 0},
	{Value: 114, Coor: []int{5, 8}, Line: 0},
}

func TestMap(t *testing.T) {
	gotNumbers, gotChar := Map("test_input")
	wantChar := CharMapTest
	wantNumbers := NumbersTest

	if !reflect.DeepEqual(gotChar, wantChar) {
		t.Errorf("want %v, got %v", wantChar, gotChar)
	}

	if !reflect.DeepEqual(gotNumbers, wantNumbers) {
		t.Errorf("want %v, got %v", wantNumbers, gotNumbers)
	}
}

func TestFindOverlap(t *testing.T) {
	got := FindOverlap(NumbersTest, CharMapTest)
	want := 467835

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
