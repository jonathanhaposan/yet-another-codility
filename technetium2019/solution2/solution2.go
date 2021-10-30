package solution

import (
	"strconv"
	"strings"
)

var (
	matrix    [][]int
	maxRow    int
	maxColumn int
)

// failed, 100% correctness 33% performance
// https://app.codility.com/demo/results/trainingNNRYKQ-X8T/
// the `isAdded` map is not reset properly, the map purpose is only for checking if the next node/cell from top or bot parent has been added
// example:
// 2 3 5
// 3 4 6
// 7 1 3
// because [0][1] and [1][0] have the same value we need to check the next node/cell and get the next higher value
// since [0][1] bottom node/cell is the same as [1][0] right cell, we dont want to add [1][1] twice
func Solution(A [][]int) string {
	matrix = A
	maxRow = len(matrix) - 1
	maxColumn = len(matrix[0]) - 1
	startNode := coordinate{}
	endNode := coordinate{row: maxRow, column: maxColumn}

	if maxRow == 0 && maxColumn == 0 {
		return startNode.getValueString()
	}

	list := []string{startNode.getValueString()}
	list = append(list, findCombination([]coordinate{getRightPos(startNode), getBottomPos(startNode)})...)
	list = append(list, endNode.getValueString())

	return strings.Join(list, "")
}

type coordinate struct {
	row    int
	column int
}

func (c coordinate) getValueInt() int {
	return matrix[c.row][c.column]
}

func (c coordinate) getValueString() string {
	return strconv.Itoa(c.getValueInt())
}

func (c coordinate) key() string {
	return strconv.Itoa(c.row) + "." + strconv.Itoa(c.column)
}

var isAdded = make(map[string]bool)

// {9, 9, 7},
// {9, 7, 2},
// {6, 9, 5},
// {9, 1, 2},
func findCombination(list []coordinate) []string {
	maxAnswerDigit := maxRow + maxColumn - 1
	var maxValue []string
	for i := 0; i < maxAnswerDigit; i++ {

		currMaxValue := 0
		nextCoordinate := []coordinate{}

		for _, coor := range list {

			if coor.getValueInt() == currMaxValue {
				nextCoordinate = append(nextCoordinate, addRightAndBottomPos(coor)...)
			} else if coor.getValueInt() > currMaxValue {
				currMaxValue = coor.getValueInt()
				isAdded = map[string]bool{}
				nextCoordinate = addRightAndBottomPos(coor)
			}
		}
		maxValue = append(maxValue, strconv.Itoa(currMaxValue))
		list = nextCoordinate
		isAdded = map[string]bool{}

	}

	return maxValue
}

func getRightPos(currPos coordinate) coordinate {
	if currPos.column == maxColumn {
		return getBottomPos(currPos)
	}

	return coordinate{
		row:    currPos.row,
		column: currPos.column + 1,
	}
}

func getBottomPos(currPos coordinate) coordinate {
	if currPos.row == maxRow {
		return getRightPos(currPos)
	}

	return coordinate{
		row:    currPos.row + 1,
		column: currPos.column,
	}
}

func addRightAndBottomPos(currPos coordinate) []coordinate {
	var nextCoordinate []coordinate
	rightPos := getRightPos(currPos)
	if !isAdded[rightPos.key()] {
		isAdded[rightPos.key()] = true
		nextCoordinate = append(nextCoordinate, rightPos)
	}

	botPos := getBottomPos(currPos)
	if !isAdded[botPos.key()] {
		isAdded[botPos.key()] = true
		nextCoordinate = append(nextCoordinate, botPos)
	}
	return nextCoordinate
}
