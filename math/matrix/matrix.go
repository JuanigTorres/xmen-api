package matrix

import "github.com/JuanigTorres/xmen-api/math/utils"

type Direction int

const (
	UP Direction = iota
	DOWN
)

type Matrix struct {
	Values [][]string
}

type IMatrix interface {
	Row(i int) []string
	Column(i int) []string
	Diagonals(d Direction) [][]string
}

func NewSquareMatrix(n int) *Matrix {
	return newMatrix(uint(n), uint(n))
}

func (matrix *Matrix) Row(r int) []string {
	return matrix.Values[r]
}

func (matrix *Matrix) Column(c int) []string {
	var column []string
	for i := range matrix.Values {
		column = append(column, matrix.Values[i][c])
	}
	return column
}

// Diagonals Returns a list of arrays for each diagonal found in the matrix crosswise.
// The transversal to be applied will depend from the direction passed.
// For example, for the following 3x3 Matrix
//
// 	1 2 3  = UP ===> [[1], [4,2], [7,5,3], [8,6], [9]]
//	4 5 6
//	7 8 9  = DOWN => [[3], [2,6], [1,5,9], [4,8], [7]]
func (matrix Matrix) Diagonals(direction Direction) [][]string {
	var diags [][]string
	var fill func(Matrix, int, int, int, int) []string

	switch direction {
	case UP:
		fill = upper
	case DOWN:
		fill = lower
	default:
		return [][]string{}
	}

	size := len(matrix.Values)
	lines := 2*size - 1 // Total of lines in the matrix
	for l := 0; l < lines; l++ {
		start := utils.Max(0, l-(size-1))   // Get the start index int of the line in the matrix. It's between [0, len(m))
		count := utils.Min(l+1, size-start) // The number of elements in a line. It's between [1, len(m))
		diags = append(diags, fill(matrix, size, l, start, count))
	}
	return diags
}

func upper(matrix Matrix, size, line, start, count int) []string {
	var diag []string
	for j := 0; j < count; j++ {
		x := utils.Min(size-1, line) - j
		y := start + j
		diag = append(diag, matrix.Values[x][y])
	}
	return diag
}

func lower(matrix Matrix, size, line, start, count int) []string {
	var diag []string
	for j := 0; j < count; j++ {
		x := start + j
		y := utils.Max((size-1)-line, 0) + j
		diag = append(diag, matrix.Values[x][y])
	}
	return diag
}

func newMatrix(n, m uint) *Matrix {
	matrix := make([][]string, n)
	for i := 0; i < int(n); i++ {
		matrix[i] = make([]string, m)
	}
	return &Matrix{matrix}
}
