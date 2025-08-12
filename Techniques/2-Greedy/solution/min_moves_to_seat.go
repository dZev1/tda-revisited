package solution

import (
	"math"
	"slices"
)

// See Leetcode 2037

func MinMovesToSeat(seats, students []int) int {
    slices.Sort(seats)
	slices.Sort(students)

	totalSteps := 0
	for i := range len(seats) {
		totalSteps += int(math.Abs(float64(seats[i]) - float64(students[i])))
	}

	return totalSteps
}