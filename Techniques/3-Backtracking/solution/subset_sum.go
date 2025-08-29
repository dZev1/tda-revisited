package solution

func SubsetSum(nums []int, target int) bool {
	var bt func(index, target int) bool
	bt = func(index, target int) bool {
		if target == 0 {
			return true
		}
		if index < 0 {
			return false
		}

		return bt(index-1, target-nums[index]) || bt(index-1, target)
	}

	return bt(len(nums)-1, target)
}
