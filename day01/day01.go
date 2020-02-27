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