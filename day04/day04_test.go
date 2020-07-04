package day04

import (
	"strings"
	"testing"
)

type testdata struct {
	input  string
	expect int
}

func TestSolvePart1(t *testing.T) {
	var tests = []testdata{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, testCase := range tests {
		result, err := Solve(strings.NewReader(testCase.input), true)

		if err != nil {
			t.Error(err)
		}
		if result != testCase.expect {
			t.Errorf("For \"%s\": expected %d, got %d",
				testCase.input,
				testCase.expect,
				result)
		}
	}
}
