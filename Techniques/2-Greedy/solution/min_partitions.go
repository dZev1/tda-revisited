package solution

// See Leetcode 1689

func MinPartitions(n string) int {
	min := 0
	for _, d := range n {
		digit := int(d - '0')
		min = max(min, digit)
	}
	return min
}
