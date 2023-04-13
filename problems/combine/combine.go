package combine

func combine(n int, k int) [][]int {
	var result [][]int
	buf := make([]int, 0, k)

	backtrack(1, n, k, buf, &result)

	return result
}

func backtrack(start, end, left int, current []int, result *[][]int) {
	if left == 0 {
		// this is necessary
		cp := make([]int, len(current))
		copy(cp, current)
		*result = append(*result, cp)
		return
	}

	for i := start; i <= end; i++ {
		backtrack(i+1, end, left-1, append(current, i), result)
	}
}
