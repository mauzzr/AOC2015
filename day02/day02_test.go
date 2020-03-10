package day02

import "testing"

type testdata struct {
	input  string
	expect int
}

func TestWrappingPaperArea(t *testing.T) {
	var tests = []testdata{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, testCase := range tests {
		l, w, h := parseLine(testCase.input)
		result := wrappingPaperArea(l, w, h)
		if result != testCase.expect {
			t.Errorf("For \"%s\": expected %d, got %d", testCase.input, testCase.expect, result)
		}
	}
}

func TestRibbonLength(t *testing.T) {
	var tests = []testdata{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, testCase := range tests {
		l, w, h := parseLine(testCase.input)
		result := ribbonLength(l, w, h)
		if result != testCase.expect {
			t.Errorf("For \"%s\": expected %d, got %d", testCase.input, testCase.expect, result)
		}
	}
}
