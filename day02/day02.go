package day02

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*

Solve1 solves the following AOC problem:

The elves are running low on
wrapping paper, and so they need to submit an order for more. They have a list
of the dimensions (length l, width w, and height h) of each present, and only
want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which
makes calculating the required wrapping paper for each gift a little easier:
find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves
also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet
of wrapping paper plus 6 square feet of slack, for a total of 58 square
feet.

 A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet
 of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

 All numbers in the elves' list are in feet. How many total square feet of
 wrapping paper should they order?

*/
func Solve1(input io.Reader) (totalSurfaceArea int, err error) {
	for {
		var line string
		_, err = fmt.Fscanln(input, &line)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
		var l, w, h int
		l, w, h, err = parseLine(line)
		if err != nil {
			return
		}
		totalSurfaceArea += wrappingPaperArea(l, w, h)
	}
	return
}

/*
Solve2 solves the following AOC 2015 problem:

The elves are also running low on ribbon. Ribbon is all the same width, so they
only have to worry about the length they need to order, which they would again
like to be exact.

The ribbon required to wrap a present is the shortest distance around its
sides, or the smallest perimeter of any one face. Each present also requires a
bow made out of ribbon as well; the feet of ribbon required for the perfect bow
is equal to the cubic feet of volume of the present. Don't ask how they tie the
bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap
the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.

A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap
the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14
feet.

How many total feet of ribbon should they order?
*/
func Solve2(input io.Reader) (totalLength int, err error) {
	for {
		var line string
		_, err = fmt.Fscanln(input, &line)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}

		var l, w, h int
		l, w, h, err = parseLine(line)
		if err != nil {
			return
		}

		totalLength += ribbonLength(l, w, h)
	}
	return
}

func ribbonLength(l int, w int, h int) int {
	s := []int{l, w, h}
	sort.Ints(s)

	return 2*s[0] + 2*s[1] + l*w*h
}

func wrappingPaperArea(l int, w int, h int) int {
	var min int

	t1 := l * w
	t2 := l * h
	t3 := w * h

	// Find the area of the smallest face
	if t1 < t2 {
		min = t1
	} else {
		min = t2
	}
	if t3 < min {
		min = t3
	}

	// Total paper area = surface area + area of smallest face
	return 2*(t1+t2+t3) + min
}

func parseLine(line string) (l int, w int, h int, err error) {
	// skip empty lines
	// this is ugly but at least it's unicode-aware
	buf := []byte(line)
	if utf8.RuneCount(buf) == 0 {
		return
	}

	subs := strings.Split(line, "x")
	if len(subs) != 3 {
		err = fmt.Errorf("Wrong number of dimensions: %d from %q", len(subs), line)
		return
	}

	var se error
	l, se = strconv.Atoi(subs[0])
	w, se = strconv.Atoi(subs[1])
	h, se = strconv.Atoi(subs[2])

	if l < 0 || w < 0 || h < 0 || se != nil {
		err = fmt.Errorf("Invalid box measurements %dx%dx%d", l, w, h)
	}
	return
}

/*
Solver holds the File reference for this day's input and implements the
SolutionPrinter interface.
*/
type Solver struct {
	Input *os.File
}

/*
PrintSolutions prints the solutions for this day's problems.
*/
func (s Solver) PrintSolutions() {
	defer s.Input.Close()
	res, err := Solve1(s.Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Day 2, Part 1: The elves should order %d square feet of wrapping paper.\n", res)
	}

	s.Input.Seek(0, io.SeekStart)

	res, err = Solve2(s.Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Day 2, Part 2: The elves should order %d feet of ribbon.\n", res)
	}
}
