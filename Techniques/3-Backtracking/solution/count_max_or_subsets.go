package solution

// See Leetcode 2044

func CountMaxOrSubsets(nums []int) int {
	maxVal := 0

	for _, num := range nums {
		maxVal |= num
	}

	var count func(index, curr int) int
	count = func(index, curr int) int {
		if curr == maxVal {
			notProcessed := len(nums) - index
			return 1 << notProcessed
		}

		if index == len(nums) {
			return 0
		}

		return count(index+1, curr|nums[index]) + count(index+1, curr)
	}

	return count(0, 0)
}
