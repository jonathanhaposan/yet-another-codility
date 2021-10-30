package solution

import (
	"fmt"
	"strconv"
	"strings"
)

// Failed, 100% correctnes 0% performance
// https://app.codility.com/demo/results/training8VKVG4-55W/
// https://app.codility.com/demo/results/trainingSSXGN7-MNE/
// https://app.codility.com/demo/results/trainingEBRES4-M3K/
// why these solution failed on performance, because this is basically a brute force solution
// if we have a large matrix and all the node/cell value is equal, we have to visit every node/cell
func Solution(A [][]int) string {
	matrix = A
	rowLength = len(A) - 1
	columnLength = len(A[0]) - 1

	startPos := coordinate{}
	asd := strings.Join(findNode(startPos), "")

	return asd
}

var (
	matrix                  [][]int
	rowLength, columnLength int
)

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

var memo = make(map[string][]string)

func findNode(currPosition coordinate) []string {
	list := []string{currPosition.getValueString()}

	if val, exist := memo[currPosition.key()]; exist {
		fmt.Println("fetching", currPosition.key())
		memo[currPosition.key()] = list
		return val
	}

	if currPosition.column == columnLength && currPosition.row == rowLength {
		memo[currPosition.key()] = list
		return list
	} else if currPosition.column == columnLength {
		list = append(list, findNode(getBottomPos(currPosition))...)
		memo[currPosition.key()] = list
		return list
	} else if currPosition.row == rowLength {
		list = append(list, findNode(getRightPos(currPosition))...)
		memo[currPosition.key()] = list
		return list
	}

	rightNeighbor := getRightPos(currPosition)
	bottomNeighbor := getBottomPos(currPosition)

	if rightNeighbor.getValueInt() == bottomNeighbor.getValueInt() {
		list = append(list, splitRecursive1(rightNeighbor, bottomNeighbor)...)
	} else if rightNeighbor.getValueInt() > bottomNeighbor.getValueInt() {
		list = append(list, findNode(rightNeighbor)...)
	} else {
		list = append(list, findNode(bottomNeighbor)...)
	}

	memo[currPosition.key()] = list
	return list
}

func splitRecursive1(rightPos, bottomPos coordinate) []string {
	var rightList []string
	rightList = append(rightList, findNode(rightPos)...)

	var bottomList []string
	bottomList = append(bottomList, findNode(bottomPos)...)

	rightListVal, _ := strconv.Atoi(strings.Join(rightList, ""))
	bottomListVal, _ := strconv.Atoi(strings.Join(bottomList, ""))

	if rightListVal > bottomListVal {
		return rightList
	}
	return bottomList

}

func splitRecursive2(top, bot coordinate) coordinate {
	topRight := getRightPos(top)
	topBottom := getBottomPos(top)

	highTop := 0
	if topRight.getValueInt() > topBottom.getValueInt() {
		highTop = topRight.getValueInt()
	} else {
		highTop = topBottom.getValueInt()
	}

	botRight := getRightPos(bot)
	botBottom := getBottomPos(bot)

	highBot := 0
	if botRight.getValueInt() > botBottom.getValueInt() {
		highBot = botRight.getValueInt()
	} else {
		highBot = botBottom.getValueInt()
	}

	if highTop > highBot {
		return top
	} else {
		return bot
	}
}

func getRightPos(currPos coordinate) coordinate {
	if currPos.column == columnLength {
		return getBottomPos(currPos)
	}

	return coordinate{
		row:    currPos.row,
		column: currPos.column + 1,
	}
}

func getBottomPos(currPos coordinate) coordinate {
	if currPos.row == rowLength {
		return getRightPos(currPos)
	}

	return coordinate{
		row:    currPos.row + 1,
		column: currPos.column,
	}
}
