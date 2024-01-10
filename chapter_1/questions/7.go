package main

// example:
// [
// 	[1, 0, 0, 0],
// 	[0, 0, 0, 0],
// 	[0, 0, 0, 0],
// 	[0, 0, 0, 0],
// ]
// [
// 	[1, 0, 0, 0, 0],
// 	[0, 0, 0, 0, 0],
// 	[0, 0, 0, 0, 0],
// 	[0, 0, 0, 0, 0],
// 	[0, 0, 0, 0, 0],
// ]

// assume size of matrix is length x length
func RotateMatrix(matrix [][]int, length int) [][]int {
	var limit int
	if length%2 == 0 {
		limit = length / 2
	} else {
		limit = length/2 + 1
	}

	for x := 0; x < limit; x++ {
		for y := 0; y < limit; y++ {
			rotated0 := matrix[x][y]
			rotated90 := matrix[y][length-1-x]
			rotated180 := matrix[length-1-x][length-1-y]
			rotated270 := matrix[length-1-y][x]

			matrix[x][y] = rotated270
			matrix[y][length-1-x] = rotated0
			matrix[length-1-x][length-1-y] = rotated90
			matrix[length-1-y][x] = rotated180
		}
	}

	return matrix
}
