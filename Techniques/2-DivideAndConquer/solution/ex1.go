package solution

func MoreLeftArr(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	sumLeft := 0
	sumRight := 0

	mid := len(arr) / 2

	for i := 0; i < mid; i++ {
		sumLeft += arr[i]
	}

	for i := mid; i < len(arr); i++ {
		sumRight += arr[i]
	}
	
	isMoreLeftL := MoreLeftArr(arr[:mid])
	isMoreLeftR := MoreLeftArr(arr[mid:])
	return (sumLeft > sumRight) && isMoreLeftL && isMoreLeftR
}