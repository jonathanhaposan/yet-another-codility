package technetium2019

var matrix [][]int

func testcase() {

	// matrix = [][]int{
	// 	{9, 9, 7},
	// 	{9, 7, 2},
	// 	{6, 9, 5},
	// 	{9, 1, 2},
	// }

	// matrix = [][]int{
	// 	{9, 5, 5, 6, 4},
	// 	{5, 9, 1, 1, 9},
	// 	{1, 9, 1, 1, 4},
	// 	{1, 1, 1, 1, 5},
	// 	{1, 1, 1, 1, 3},
	// }

	size := 999

	for i := 0; i < size; i++ {
		temp := []int{}
		for j := 0; j < size; j++ {
			temp = append(temp, 9)
		}
		matrix = append(matrix, temp)
	}
}
