package solution

func SubsetXORSum(nums []int) int {
	var backtracking func(index, sum int) int
	backtracking = func(index, sum int) int {
		if index == len(nums) {
			return sum
		}
		return backtracking(index + 1, sum ^ nums[index]) + backtracking(index + 1, sum)
	}

	return backtracking(0, 0)
}