package slices

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Predicate represents a predicate (boolean-value function) of one argument
type Predicate[T any] func(t T) bool

// Filter returns a new slice containing all the elements that satisfied the
// Predicate.
func Filter[T any](s []T, fn Predicate[T]) []T {
	res := make([]T, 0)
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}

// FindFirst returns the first element in a slice that satisfies the Predicate
// and a boolean indicating if found. Once an element in the slice satisfied the
// Predicate it stops processing elements.
func FindFirst[T any](slice []T, fn Predicate[T]) (res T, ok bool) {
	for i := 0; i < len(slice); i++ {
		if fn(slice[i]) {
			return slice[i], true
		}
	}
	return res, false
}

// FindAll returns a slice containers all the elements for which the Predicate is
// satisfied. If no elements satisfy the Predicate an empty slice is returned.
func FindAll[T any](slice []T, fn Predicate[T]) []T {
	results := make([]T, 0)
	for i := 0; i < len(slice); i++ {
		if fn(slice[i]) {
			results = append(results, slice[i])
		}
	}
	return results
}

// Contains returns true if the slice contains at least one occurrence of the
// specified element.
func Contains[T comparable](slice []T, item T) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return true
		}
	}
	return false
}

// Count returns the number of occurrences item is found in the slice.
func Count[T comparable](slice []T, item T) int {
	count := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			count++
		}
	}
	return count
}

// Remove will remove all instances of a given element from the slice and return
// the count of items removed.
func Remove[T comparable](slice []T, item T) ([]T, int) {
	removed := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			removed++
			slice = append(slice[:i], slice[i+1:]...)
			i-- // since slice[i] was removed that index must be reprocessed
		}
	}
	return slice, removed
}

// Map creates a new slice mapping the values that result from applying the
// map function.
func Map[T, R any](slice []T, fn func(item T) R) []R {
	results := make([]R, 0)
	for i := 0; i < len(slice); i++ {
		results = append(results, fn(slice[i]))
	}
	return results
}

// Reverse reverses the elements of a slice/array in place.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Shuffle accepts a slice and shuffles the elements of the slice randomly
// in place.
func Shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// Batch accepts a slice and a batch size returning the subset of the original slice
// according to the batch size provided.
//
// Batch can be useful when processing large volumes of data and needing to batch it
// in chunks for performance reasons.
//
// Providing a batch size less than 1 will result in a panic.
func Batch[T any](slice []T, batchSize int) [][]T {
	if batchSize < 1 {
		panic("illegal batchSize, batch size cannot be less than 1")
	}
	batches := make([][]T, 0)
	for i := 0; i < len(slice); i += batchSize {
		end := i + batchSize
		if end > len(slice) {
			end = len(slice)
		}
		batches = append(batches, slice[i:end])
	}
	return batches
}

// Equal compares two slices to determine if they are equal. Slices are considered
// equals if their lengths are the same and each element is the same, IE order
// matters.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// Clone creates a new slice and copies the contents of the provided slice into
// the returned slice. If slice is nil then nil is returned to preserve nil.
func Clone[T any](slice []T) []T {
	if slice == nil {
		return nil
	}
	cloned := make([]T, len(slice))
	copy(cloned, slice)
	return cloned
}

// Index returns the index of the first occurrence of item found in the slice.
// If the item wasn't found in the slice -1 is returned.
func Index[T comparable](slice []T, item T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Insert inserts an item at the given index of the slice and returns the modified
// slice. If the index is out of bounds this will panic.
func Insert[T any](slice []T, idx int, item T) []T {
	tot := len(slice) + 1
	if tot <= cap(slice) {
		s2 := slice[:tot]
		copy(s2[idx+1:], slice[idx:])
		copy(s2[idx:], []T{item})
		return s2
	}
	s2 := make([]T, tot)
	copy(s2, slice[:idx])
	copy(s2[idx:], []T{item})
	copy(s2[idx+1:], slice[idx:])
	return s2
}
