package day04

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const numThreads int = 4

/*

Solve solves the following AOC 2015 problem. Part 1 if solvePart1 is true, or
Part 2 otherwise:

=== PART 1 ===
Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of
abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest
such number to do so.

If your secret key is pqrstuv, the lowest number it combines with to make an
MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of
pqrstuv1048970 looks like 000006136ef....

=== PART 2 ===
Now find one that starts with six zeroes.

*/
func Solve(in io.Reader, solvePart1 bool) (suffix int, err error) {
	var prefix string
	_, err = fmt.Fscanln(in, &prefix)
	if err != nil {
		return 0, err
	}

	var mask byte = 0xFF
	if solvePart1 {
		mask = 0xF0
	}

	results := make(chan int)

	for i := 1; i <= numThreads; i++ {
		go md5Crack(prefix, i, mask, results)
	}

	// TODO: Race condition here, but it's Saturday so
	//  I don't happen to care at the moment
	select {
	case suffix = <-results:
		return suffix, err
	case <-time.After(5 * time.Second):
		return 0, fmt.Errorf("Timed out waiting for result")
	}

}

func md5Crack(prefix string, n int, lastMask byte, result chan<- int) {
	for {
		suffix := strconv.Itoa(n)
		data := prefix + suffix

		hash := md5.Sum([]byte(data))

		if !(hash[0]&0xFF > 0 || hash[1]&0xFF > 0 || hash[2]&lastMask > 0) {
			result <- n
			break
		}
		n += numThreads
	}
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
	res, err := Solve(s.Input, true)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Day 4, Part 1: Suffix is %d\n", res)
	}

	s.Input.Seek(0, io.SeekStart)

	res, err = Solve(s.Input, false)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Day 4, Part 2: Suffix is %d\n", res)
	}
}
