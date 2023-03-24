package problems

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = nums1[:m]
	nums2 = nums2[:n]

	nums1 = append(nums1, nums2[:]...)
	for i, _ := range nums1 {
		for j, _ := range nums1 {
			if nums1[i] < nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}

	}

	return nums1

}
