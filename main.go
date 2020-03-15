package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mauzzr/AOC2015/day01"
	"github.com/mauzzr/AOC2015/day02"
	"github.com/mauzzr/AOC2015/day03"
)

func openFile(path string) (f *os.File) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	var res interface{}
	var err error

	day01Input := openFile("input/day01.txt")
	defer day01Input.Close()
	res, err = day01.Solve1(day01Input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 1, Part 1: Santa ends on floor %d\n", res)
	res, err = day01.Solve2(day01Input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 1, Part 2: Santa enters the basement after %d steps.\n", res)

	day02Input := openFile("input/day02.txt")
	defer day02Input.Close()
	res, err = day02.Solve1(day02Input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 2, Part 1: The elves should order %d square feet of wrapping paper.\n", res)
	day02Input.Seek(0, io.SeekStart)
	res, err = day02.Solve2(day02Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Printf("Day 2, Part 2: The elves should order %d feet of ribbon.\n", res)

	day03Input := openFile("input/day03.txt")
	defer day03Input.Close()
	res, err = day03.Solve1(day03Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Printf("Day 3, Part 1: %d houses receive at least one present\n", res)
	day03Input.Seek(0, io.SeekStart)
	res, err = day03.Solve2(day03Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Printf("Day 3, Part 2: %d houses receive at least one present.\n", res)
}
