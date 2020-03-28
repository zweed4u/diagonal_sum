package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int) int {
	numRows := len(arr)
	numCols := len(arr[0])
	// rows should be equal to cols
	if numRows != numCols {
		fmt.Println("Number of rows does not match number of columns for each row in matrix")
		return -1
	}
	primaryDiagonal := 0   // l to right
	secondaryDiagonal := 0 // r to left
	for index := 0; index < numRows; index++ {
		primaryDiagonal += arr[index][index]
		secondaryDiagonal += arr[index][numCols-1-index]
	}
	total := float64(primaryDiagonal - secondaryDiagonal)
	return int(math.Abs(total))
}

func main() {
	/*
		|(1+10+22) - (6+10+18)|
	*/
	arr := [][]int{
		{1, 4, 6},
		{8, 10, 12},
		{18, 20, 22},
	}
	r := diagonalDifference(arr)
	fmt.Println(r)
}
