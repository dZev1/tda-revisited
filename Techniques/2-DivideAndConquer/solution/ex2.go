package solution

func MirrorIndex(arr []int) bool {
	return binarySearch(arr, 0, len(arr)-1)
}

func binarySearch(arr []int, left, right int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	numToCompare := mid + 1
	if arr[mid] == numToCompare {
		return true
	} else if arr[mid] < numToCompare {
		return binarySearch(arr, mid+1, right)
	} else {
		return binarySearch(arr, left, mid-1)
	}
}
