package solution

import "math"

func ParejasDeBaile(b1, b2 []int) int {
	i := 0
	j := 0
	cantParejas := 0

	for i < len(b1) && j < len(b2) {
		if math.Abs(float64(b1[i])-float64(b2[j])) <= 1 {
			cantParejas++
			i++
			j++
		} else if b1[i] < b2[j] {
			i++
		} else if b1[i] > b2[j] {
			j++
		}
	}
	// O(N + M)

	return cantParejas
}
