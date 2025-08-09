package solution

import (
	"greedy/hashset"
)

func SumaSelectiva(X hashset.HashSet[int], k int) (hashset.HashSet[int], int) {
	S := hashset.New[int]()
	sum := 0

	for range k {
		a, _ := X.Max()
		sum += a
		S.Add(a)
		X.Delete(a)
	}

	return *S, sum
}
