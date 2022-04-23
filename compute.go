package powerset

import "math"

// Compute returns the powerset of the given set.
func Compute[T any](set []T) [][]T {
	cap := int(math.Pow(2, float64(len(set))) - 1)
	res := make([][]T, 0, cap)
	l := len(set)
	j := 0
	for i := 0; i <= cap; i++ {
		c := make([]T, 0, 1)
		for j = 0; j < l; j++ {
			if i&(1<<j) > 0 {
				c = append(c, set[j])
			}
		}
		res = append(res, c)
	}

	return res
}
