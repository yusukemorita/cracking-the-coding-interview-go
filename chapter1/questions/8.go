package main

import "fmt"

// if an element in the MxN matrix is 0,
// the whole row / column is set to 0
// this uses O(M+N) space for storing which rows and columns
// have 0
func ZeroMatrix1(matrix [][]int) [][]int {
	zeroColumns := map[int]bool{}
	zeroRows := map[int]bool{}

	for rowIndex, row := range matrix {
		for columnIndex, val := range row {
			if val == 0 {
				zeroColumns[columnIndex] = true
				zeroRows[rowIndex] = true
			}
		}
	}

	for rowIndex, row := range matrix {
		for columnIndex, _ := range row {
			_, rowHasZero := zeroRows[rowIndex]
			_, columnHasZero := zeroColumns[columnIndex]

			if rowHasZero || columnHasZero {
				matrix[rowIndex][columnIndex] = 0
			}
		}
	}

	return matrix
}

// store the information about which rows/columns have zeros
// inside the matrix itself, specifically the first row and first column
// in order to use no additional space
func ZeroMatrix2(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	firstRowHasZero := false
	for _, value := range matrix[0] {
		if value == 0 {
			firstRowHasZero = true
			break
		}
	}

	firstColumnHasZero := false
	for _, row := range matrix {
		if row[0] == 0 {
			firstColumnHasZero = true
			break
		}
	}

	// find all zeroes in the rest of the matrix
	for rowIndex, row := range matrix {
		for columnIndex, val := range row {
			if rowIndex == 0 || columnIndex == 0 {
				continue
			}

			if val == 0 {
				// mark row as having a zero
				matrix[rowIndex][0] = 0
				// mark column as having a zero
				matrix[0][columnIndex] = 0
			}
		}
	}

	fmt.Println("first row and column marked with zero")
	printMatrix(matrix)

	// find all columns that have a zero and replace all values with zero
	for columnIndex, val := range matrix[0] {
		if columnIndex == 0 {
			continue
		}

		if val == 0 {
			fmt.Printf("column %d has a zero\n", columnIndex)
			zeroifyColumn(matrix, columnIndex)
		}
	}

	// find all rows that have a zero and replace all values with zero
	for rowIndex, row := range matrix {
		if rowIndex == 0 {
			continue
		}

		fmt.Print(row)
		fmt.Println(row)
		if row[0] == 0 {
			fmt.Printf("row %d has a zero\n", rowIndex)
			zeroifyRow(matrix, rowIndex)
		}
	}

	if firstColumnHasZero {
		zeroifyColumn(matrix, 0)
	}

	if firstRowHasZero {
		zeroifyRow(matrix, 0)
	}

	return matrix
}

func zeroifyColumn(matrix [][]int, columnIndex int) {
	for rowIndex, _ := range matrix {
		matrix[rowIndex][columnIndex] = 0
	}
}

func zeroifyRow(matrix [][]int, rowIndex int) {
	for columnIndex, _ := range matrix[rowIndex] {
		matrix[rowIndex][columnIndex] = 0
	}
}

func printMatrix(matrix [][]int) {
	fmt.Println("")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("")
}

