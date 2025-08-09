package hashset

import (
	"cmp"
	"iter"
)

// HashSet is a generic set based on a hash table (map) with max and min tracking.
type HashSet[T cmp.Ordered] struct {
	m            map[T]struct{}
	max          T
	min          T
	initialized  bool
}

// New creates a new HashSet.
func New[T cmp.Ordered]() *HashSet[T] {
	return &HashSet[T]{m: make(map[T]struct{})}
}

// InitWith creates a new HashSet initialized with vals.
func InitWith[T cmp.Ordered](vals ...T) *HashSet[T] {
	hs := New[T]()
	if len(vals) > 0 {
		hs.initialized = true
		hs.min = vals[0]
		hs.max = vals[0]
	}
	for _, v := range vals {
		hs.Add(v)
	}
	return hs
}

// Add adds a value to the set.
func (hs *HashSet[T]) Add(val T) {
	if !hs.initialized {
		hs.min = val
		hs.max = val
		hs.initialized = true
	} else {
		if val < hs.min {
			hs.min = val
		}
		if val > hs.max {
			hs.max = val
		}
	}
	hs.m[val] = struct{}{}
}

// Contains reports whether the set contains the given value.
func (hs *HashSet[T]) Contains(val T) bool {
	_, ok := hs.m[val]
	return ok
}

// Len returns the size/length of the set.
func (hs *HashSet[T]) Len() int {
	return len(hs.m)
}

// Delete removes a value from the set.
func (hs *HashSet[T]) Delete(val T) {
	if !hs.Contains(val) {
		return
	}
	
	delete(hs.m, val)
	
	if len(hs.m) == 0 {
		hs.initialized = false
		return
	}
	
	// Only need to recalculate if we deleted the current min or max
	if val == hs.max || val == hs.min {
		hs.recalculateMinMax()
	}
}

// recalculateMinMax finds new min and max values after deletion
func (hs *HashSet[T]) recalculateMinMax() {
	if len(hs.m) == 0 {
		hs.initialized = false
		return
	}
	
	first := true
	for v := range hs.m {
		if first {
			hs.min = v
			hs.max = v
			first = false
			continue
		}
		
		if v < hs.min {
			hs.min = v
		}
		if v > hs.max {
			hs.max = v
		}
	}
}

// All returns an iterator over all the values in the set.
func (hs *HashSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range hs.m {
			if !yield(v) {
				return
			}
		}
	}
}

// Max returns the maximum value in the set in O(1) time.
// Returns zero value and false if the set is empty.
func (hs *HashSet[T]) Max() (T, bool) {
	if !hs.initialized {
		var zero T
		return zero, false
	}
	return hs.max, true
}

// Min returns the minimum value in the set in O(1) time.
// Returns zero value and false if the set is empty.
func (hs *HashSet[T]) Min() (T, bool) {
	if !hs.initialized {
		var zero T
		return zero, false
	}
	return hs.min, true
}

// Union returns the set union of hs with other.
func (hs *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		result.Add(v)
	}
	for v := range other.m {
		result.Add(v)
	}
	return result
}

// Intersection returns the set intersection of hs with other.
func (hs *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns the set difference hs - other.
func (hs *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Equals compares two HashSets and returns true if they contain the same elements.
func (hs *HashSet[T]) Equals(other *HashSet[T]) bool {
    // If lengths are different, they can't be equal
    if hs.Len() != other.Len() {
        return false
    }

    // Check if every element in hs exists in other
    for v := range hs.m {
        if !other.Contains(v) {
            return false
        }
    }

    return true
}