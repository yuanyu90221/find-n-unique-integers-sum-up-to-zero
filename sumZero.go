package sum_zero

func sumZero(n int) []int {
	result := make([]int, n)
	M := n / 2
	if n%2 == 0 {
		for idx := range result {
			if idx < M {
				result[idx] = idx + 1
			} else {
				result[idx] = M - idx - 1
			}
		}
	} else {
		for idx := range result {
			if idx < M {
				result[idx] = idx + 1
			} else if idx > M {
				result[idx] = M - idx
			} else {
				result[idx] = 0
			}
		}
	}
	return result
}
