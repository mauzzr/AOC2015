package main

import (
	"os"

	"github.com/mauzzr/AOC2015/day01"
	"github.com/mauzzr/AOC2015/day02"
	"github.com/mauzzr/AOC2015/day03"
	"github.com/mauzzr/AOC2015/day04"
	"github.com/mauzzr/AOC2015/util"
)

func openFile(path string) (f *os.File) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	var printers = []util.SolutionPrinter{
		day01.Solver{Input: openFile("input/day01.txt")},
		day02.Solver{Input: openFile("input/day02.txt")},
		day03.Solver{Input: openFile("input/day03.txt")},
		day04.Solver{Input: openFile("input/day04.txt")},
	}

	for _, printer := range printers {
		printer.PrintSolutions()
	}
}
