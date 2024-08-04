package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(transposeMatrix(inputMatrix()))
}

func inputMatrix() [][]int {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func transposeMatrix(matrix [][]int) [][]int {
	t := time.Now()
	transposedMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(transposedMatrix); i++ {
		transposedMatrix[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			transposedMatrix[i][j] = matrix[j][i]
		}
	}
	fmt.Println(time.Since(t))
	return transposedMatrix
}
