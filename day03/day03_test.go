package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart01(t *testing.T) {
	result := Part01(testInput)
	expected := 157

	if result != expected {
		t.Errorf("Part01 value is incorrect. Expected = %d, Result = %d", expected, result)
	}
}

func TestPart02(t *testing.T) {
	result := Part02(testInput)
	expected := 70

	if result != expected {
		t.Errorf("Part02 value is incorrect. Expected = %d, Result = %d", expected, result)
	}
}
