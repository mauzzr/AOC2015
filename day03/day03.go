package day03

import (
	"fmt"
	"io"
	"os"
)

type point struct {
	x int
	y int
}

/*

Solve1 solves the following AOC 2015 problem:

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and
then an elf at the North Pole calls him via radio and tells him where to move
next. Moves are always exactly one house to the north (^), south (v), east (>),
or west (<). After each move, he delivers another present to the house at his
new location.

However, the elf back at the north pole has had a little too much eggnog, and
so his directions are a little off, and Santa ends up visiting some houses more
than once. How many houses receive at least one present?

For example:

	> delivers presents to 2 houses: one at the starting location, and one to
	the east.

	^>v< delivers presents to 4 houses in a square, including twice to the house
	at his starting/ending location.

 ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2
 houses.

*/
func Solve1(input io.Reader) (uniqueHouses int, err error) {
	var currentLocation point
	houses := make(map[point]int)

	houses[currentLocation] = 1

	for {
		var r rune
		_, se := fmt.Fscanf(input, "%c", &r)
		if se == io.EOF {
			break
		} else if se != nil {
			err = se
			return
		}

		switch r {
		case '>':
			currentLocation.x++
		case '<':
			currentLocation.x--
		case '^':
			currentLocation.y++
		case 'v', 'V':
			currentLocation.y--
		default:
			err = fmt.Errorf("Unexpected character %c", r)
			return
		}

		houses[currentLocation]++
	}

	uniqueHouses = len(houses)

	return
}

/*

Solve2 solves the following AOC 2015 problem:

The next year, to speed up the process, Santa creates a robot version of
himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the
same starting house), then take turns moving based on instructions from the
elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

 ^v delivers presents to 3 houses, because Santa goes north, and then
 Robo-Santa goes south.

 ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back
 where they started.

 ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction
 and Robo-Santa going the other.

*/
func Solve2(input io.Reader) (uniqueHouses int, err error) {
	var santaLoc, roboSantaLoc point
	var i int
	houses := make(map[point]int)
	houses[point{0, 0}] = 2

	for {
		var r rune
		_, se := fmt.Fscanf(input, "%c", &r)
		if se == io.EOF {
			break
		} else if se != nil {
			err = se
			return
		}

		var currentLocation *point
		if (i % 2) == 0 {
			currentLocation = &santaLoc
		} else {
			currentLocation = &roboSantaLoc
		}

		switch r {
		case '>':
			currentLocation.x++
		case '<':
			currentLocation.x--
		case '^':
			currentLocation.y++
		case 'v', 'V':
			currentLocation.y--
		default:
			err = fmt.Errorf("Unexpected character %c", r)
			return
		}

		houses[*currentLocation]++
		i++
	}

	uniqueHouses = len(houses)
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
		fmt.Printf("Day 3, Part 1: %d houses receive at least one present\n", res)
	}

	s.Input.Seek(0, io.SeekStart)

	res, err = Solve2(s.Input)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Day 3, Part 2: %d houses receive at least one present.\n", res)
	}
}
