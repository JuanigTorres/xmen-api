package utils

import "math"

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// SecureDiv makes a division, if the divisor is 0 the result is 0 too.
func SecureDiv(a, b int64) float32 {
	if b == 0 {
		return float32(0)
	}
	return float32(a) / float32(b)
}
