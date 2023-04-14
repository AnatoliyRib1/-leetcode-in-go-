package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	buf := make([]int, 0, len(candidates))

	backtrack(candidates, buf, target, &result)
	return result
}

func backtrack(candidates []int, buf []int, left int, result *[][]int) {

	if left == 0 {
		cp := make([]int, len(buf))
		copy(cp, buf)
		*result = append(*result, cp)
		return
	}

	for i, candidate := range candidates {
		if candidate <= left {
			backtrack(candidates[i:], append(buf, candidate), left-candidate, result)
		}
	}
	return
}
