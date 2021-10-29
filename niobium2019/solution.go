package solution

import "strings"

// Accepted
// https://app.codility.com/demo/results/trainingVYTA5U-AUX/
// The logic is to find the most repeated pattern from the matrix
// By taking the first value from each row, we can make the pattern like this
// Example:
// 0 1 1 1	-> 	ok flip flip flip
// 0 1 1 1	   	ok flip flip flip
// 0 0 0 1 	   	ok  ok   ok  flup
// 0 1 1 1		ok flip flip flip
//
// Notice that `ok flip flip flip` is shown multiple times, so return 3 as the solution
// By look the input matrix the ideal flip should be happen on the first column only
// But the problem only need to find the maximum number of rows that containing same value
// Not how efficient flipping the column, hence the solution works
func Solution1(A [][]int) int {

	countMap := make(map[string]int)

	for _, row := range A {

		var temp = []string{}
		for _, column := range row {
			if row[0] == column {
				temp = append(temp, "ok")
			} else {
				temp = append(temp, "flip")
			}
		}

		pattern := strings.Join(temp, ".")
		countMap[pattern]++
	}

	highScore := 0
	for _, score := range countMap {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}
