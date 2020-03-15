package day02

import (
	"strings"
	"testing"
)

type testdata struct {
	input  string
	expect int
}

var wrappingPaperAreaTests = []testdata{
	{"2x3x4", 58},
	{"1x1x10", 43},
}

var ribbonLengthTests = []testdata{
	{"2x3x4", 34},
	{"1x1x10", 14},
}

func TestWrappingPaperArea(t *testing.T) {
	for _, testCase := range wrappingPaperAreaTests {
		l, w, h, err := parseLine(testCase.input)
		if err != nil {
			t.Error(err)
		}

		result := wrappingPaperArea(l, w, h)
		if result != testCase.expect {
			t.Errorf("For \"%s\": expected %d, got %d", testCase.input, testCase.expect, result)
		}
	}
}

func TestRibbonLength(t *testing.T) {
	for _, testCase := range ribbonLengthTests {
		l, w, h, err := parseLine(testCase.input)
		if err != nil {
			t.Error(err)
		}

		result := ribbonLength(l, w, h)
		if result != testCase.expect {
			t.Errorf("For \"%s\": expected %d, got %d", testCase.input, testCase.expect, result)
		}
	}
}

func streamifyTestData(testCases []testdata) *strings.Reader {
	inputs := make([]string, len(testCases))

	for i, testCase := range testCases {
		inputs[i] = testCase.input
	}

	lines := strings.Join(inputs, "\n")

	return strings.NewReader(lines)
}

func TestSolve1(t *testing.T) {
	var expect int
	input := streamifyTestData(wrappingPaperAreaTests)

	for _, testCase := range wrappingPaperAreaTests {
		expect += testCase.expect
	}
	res, err := Solve1(input)
	if err != nil {
		t.Error(err)
	}
	if res != expect {
		t.Errorf("In Solve1: Expected %d, got %d", expect, res)
	}
}

func TestSolve2(t *testing.T) {
	var expect int
	input := streamifyTestData(ribbonLengthTests)

	for _, testCase := range ribbonLengthTests {
		expect += testCase.expect
	}

	res, err := Solve2(input)
	if err != nil {
		t.Error(err)
	}
	if res != expect {
		t.Errorf("In Solve2: Expected %d, got %d", expect, res)
	}
}
