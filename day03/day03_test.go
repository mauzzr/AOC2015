package day03

import (
	"strings"
	"testing"
)

type testData struct {
	input  string
	expect int
}

func TestSolve1(t *testing.T) {
	var testCases = []testData{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, testCase := range testCases {
		res, err := Solve1(strings.NewReader(testCase.input))
		if err != nil {
			t.Error(err)
		}
		if res != testCase.expect {
			t.Errorf("For %q: expected %d, got %d", testCase.input, testCase.expect, res)
		}
	}
}

func TestSolve2(t *testing.T) {
	var testCases = []testData{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, testCase := range testCases {
		res, err := Solve2(strings.NewReader(testCase.input))
		if err != nil {
			t.Error(err)
		}
		if res != testCase.expect {
			t.Errorf("For %q: expected %d, got %d", testCase.input, testCase.expect, res)
		}
	}
}
