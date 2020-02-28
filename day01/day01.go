package day01

/*

Solve1 solves the following AOC problem:

Santa is trying to deliver presents in a large apartment building, but he
can't find the right floor - the directions he got are a little confusing. He
starts on the ground floor (floor 0) and then follows the instructions one
character at a time.

An opening parenthesis, (, means he should go up one floor, and a closing
parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will
never find the top or bottom floors.

*/
func Solve1(input string) (floor int) {
	for _, r := range input {
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}
	}
	return
}

/*

Solve2 solves the following AOC problem:

Now, given the same instructions, find the position of the first character that
causes him to enter the basement (floor -1). The first character in the
instructions has position 1, the second character has position 2, and so on.

For example:

    ) causes him to enter the basement at character position 1.
    ()()) causes him to enter the basement at character position 5.

What is the position of the character that causes Santa to first enter the
basement?

*/
func Solve2(input string) (step int) {
	floor := 0
	for i, r := range input {
		step = i + 1
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}
		if floor < 0 {
			return
		}
	}
	return
}
