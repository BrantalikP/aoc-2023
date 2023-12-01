package main

import (
	"reflect"
	"testing"
)

const (
	validValue   = "1abc2"
	invalidValue = "asd"
	singleValue  = "asds1"
)

var exampleValue = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func TestGetCalibration(t *testing.T) {
	t.Run("Get Calibration values", func(t *testing.T) {
		got := GetCalibrationValue(validValue)
		want := 12

		if got != want {
			t.Errorf("want %v but got %v", got, want)
		}
	})

	t.Run("Get Calibration values from invalid", func(t *testing.T) {
		got := GetCalibrationValue(invalidValue)
		want := 0

		if got != want {
			t.Errorf("want %v but got %v", got, want)
		}
	})

	t.Run("Get Calibration values from single ", func(t *testing.T) {
		got := GetCalibrationValue(singleValue)
		want := 11

		if got != want {
			t.Errorf("want %v but got %v", got, want)
		}
	})
}

func TestGetSumOfCalibrationValues(t *testing.T) {
	got := GetSumOfCalibrationValues(exampleValue)
	want := 142

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, exampleValue)
	}
}

func TestScanAndReturn(t *testing.T) {
	got := ScanAndReturn("testValues")
	want := exampleValue

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestGetResult(t *testing.T) {
	got := GetResult("testValuesWords")
	want := 281
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
}

func TestReplaceWrittenNumbers(t *testing.T) {
	t.Run("text numbers are replaced correctly", func(t *testing.T) {
		got := ReplaceWrittenNumbers("abcone2threexyz")
		want := "abc123xyz"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("text numbers are replaced correctly - string without numbers", func(t *testing.T) {
		got := ReplaceWrittenNumbers("abc2texyz")
		want := "abc2texyz"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
