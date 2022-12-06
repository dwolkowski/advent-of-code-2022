package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart01(t *testing.T) {
	result := Part01(testInput)
	expected := "CMZ"

	if result != expected {
		t.Errorf("Part01 value is incorrect. Expected = %s, Result = %s", expected, result)
	}
}

func TestPart02(t *testing.T) {
	result := Part02(testInput)
	expected := ""

	if result != expected {
		t.Errorf("Part02 value is incorrect. Expected = %s, Result = %s", expected, result)
	}
}
