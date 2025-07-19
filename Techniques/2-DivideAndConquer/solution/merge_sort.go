package solution


func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr) / 2
	left := MergeSort(arr[:n])
	right := MergeSort(arr[n:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}