package main

import (
	"fmt"
	"os"

	aoc "github.com/golang-insiders/advent-of-code-2023/library"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input file path not provided")
		return
	}
	input := aoc.ReadFileLineByLine(os.Args[1])
	matrices := generateMartices(input)
	fmt.Println("score", solveMatrices(matrices))
}

type Matrix []string

func (m Matrix) solveMatrix() (int, int) {
	var horizontalLine, verticalLine int
	horizontalLine = m.findReflectionLine()

	if horizontalLine != 0 {
		return horizontalLine, verticalLine
	}

	m = m.transformMatrix()
	verticalLine = m.findReflectionLine()

	return horizontalLine, verticalLine
}

func (m Matrix) findReflectionLine() int {
	var n = len(m)
	for row := 0; row < n-1; row++ {
		nextRow := row + 1
		if m[row] == m[nextRow] {

			var i, j = row, nextRow
			var isReflectionPoint = true
			for i >= 0 && j < n {
				if m[i] != m[j] {
					isReflectionPoint = false
					break
				}
				i--
				j++
			}

			if isReflectionPoint {
				return row + 1
			}
		}
	}
	return 0
}

func (m Matrix) transformMatrix() []string {
	// considering matrix to be n*m, where denotes 'N' no. of rows , 'M' denotes number of column
	var N = len(m)
	var M = len(m[0])

	var transformedMatrix []string
	for i := 0; i < M; i++ {
		var str string
		for j := 0; j < N; j++ {
			str += string(m[j][i])
		}
		transformedMatrix = append(transformedMatrix, str)
	}
	return transformedMatrix
}

func generateMartices(input []string) []Matrix {
	var matrices = make([]Matrix, 0)
	var matrix Matrix
	for _, str := range input {
		if len(str) != 0 {
			matrix = append(matrix, str)
		} else {
			matrices = append(matrices, matrix)
			matrix = Matrix{}
		}
	}
	matrices = append(matrices, matrix)
	return matrices
}

func solveMatrices(matrices []Matrix) int {
	var score = 0
	var hTotal, vTotal int
	for _, matrix := range matrices {
		h, v := matrix.solveMatrix()
		hTotal += h
		vTotal += v
	}
	score = vTotal + 100*hTotal
	return score
}
