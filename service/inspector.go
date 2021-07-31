package service

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	exception "github.com/JuanigTorres/xmen-finder/exception"
	matrixutils "github.com/JuanigTorres/xmen-finder/math/matrix"
)

const NITROGEN_BASE = "ATCG"

type View uint

const (
	HORIZONTAL View = iota
	VERTICAL
	HIGHER_DIAGONAL
	LOWER_DIAGONAL
)

func IsMutant(dna []string) (bool, error) {
	matrix, err := validateAndFillAsMatrix(dna)

	if err != nil {
		log.Println("ERROR :: ", err)
		return false, err
	}

	dimensions := map[View] [][]string {
		HORIZONTAL	   : {},
		VERTICAL	   : {},
		HIGHER_DIAGONAL: {},
		LOWER_DIAGONAL : {},
	}

	dimensions[HIGHER_DIAGONAL] = matrix.Diagonals(matrixutils.UP)
	dimensions[LOWER_DIAGONAL] 	= matrix.Diagonals(matrixutils.DOWN)
	for i := range matrix.Values {
		dimensions[HORIZONTAL] = append(dimensions[HORIZONTAL],  matrix.Row(i))
		dimensions[VERTICAL]   = append(dimensions[VERTICAL], matrix.Column(i))
	}

	var founded int
	for _, dimension := range dimensions {
		for _, count := range consecutiveOccurrences(dimension) {
			if count >= 4 {
				founded ++
			}
		}
	}
	return founded > 1, nil
}

func validateAndFillAsMatrix(dna []string) (matrix *matrixutils.Matrix, err error) {
	n := len(dna)
	matrix = matrixutils.NewSquareMatrix(n)
	for i, chain := range dna {
		if err = isValidChain(n, chain); err != nil {
			return nil, err
		}

		for j, char := range chain {
			matrix.Values[i][j] = string(char)
		}
	}
	return matrix, nil
}

func isValidChain(size int, chain string) error {
	if size <= 0 {
		return exception.MatrixError("The size must be higher than zero.")
	}

	if len(chain) != size {
		return exception.MatrixError("The dna must be an NxN matrix.")
	}

	for _, char := range chain {
		if !unicode.IsLetter(char) || !strings.ContainsRune(NITROGEN_BASE, char) {
			msg := fmt.Sprintf("The dna with the chain [%v] doesn't contains a valid nitrogen base.", chain)
			return exception.MatrixError(msg)
		}
	}

	return nil
}

// consecutiveOccurrences return the highest number of consecutive occurrences from each Vector in the list
func consecutiveOccurrences(vectors [][]string) []int {
	var result []int
	for _, v := range vectors {
		count := 1
		best  := 0
		for i := 1; i < len(v); i++ {
			if v[i] == v[i - 1] {
				count ++
			} else {
				count = 1
			}
			if best < count {
				best = count
			}
		}
		result = append(result, best)
	}
	return result
}
