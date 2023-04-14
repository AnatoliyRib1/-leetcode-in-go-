package subsets

func subsets(nums []int) [][]int {
	var result [][]int
	buf := make([]int, 0, len(nums))
	backtrack(nums, buf, &result)
	return result
}
func backtrack(nums []int, buf []int, result *[][]int) {
	cp := make([]int, len(buf))
	copy(cp, buf)
	*result = append(*result, cp)
	for i, l := range nums {

		backtrack(nums[i+1:], append(buf, l), result) // no need to append and then remove
	}
}
