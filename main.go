package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mauzzr/AOC2015/day01"
	"github.com/mauzzr/AOC2015/day02"
)

func openFile(path string) (f *os.File) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	day01Input := openFile("input/day01.txt")
	defer day01Input.Close()
	res, err := day01.Solve1(day01Input)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Day 1, Part 1: Santa ends on floor %d\n", res)
	res, err = day01.Solve2(day01Input)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Day 1, Part 2: Santa enters the basement after %d steps.\n", res)

	f := openFile("input/day02.txt")
	var day02Input []string
	var day02Line string

	for {
		_, err := fmt.Fscanln(f, &day02Line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		day02Input = append(day02Input, day02Line)
	}
	fmt.Printf("DEBUG: %d lines\n", len(day02Input))
	fmt.Printf("Day 2, Part 1: The elves should order %d square feet of wrapping paper.\n", day02.Solve1(day02Input))
	fmt.Printf("Day 2, Part 2: The elves should order %d feet of ribbon.\n", day02.Solve2(day02Input))
}
