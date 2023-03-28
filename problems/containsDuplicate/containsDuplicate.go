package containsDuplicate

func containsDuplicate(nums []int) bool {
	myMap := map[int]int{}
	for i, value := range nums {
		myMap[value] = i
	}
	if len(nums) > len(myMap) {
		return true
	} else {
		return false
	}

}
