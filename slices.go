package slices

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

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

// TakeIf iterates through a slice and all elements that satisfy the predicate
// are passed to a function/closure to be processed. TakeIf is similar to Filter
// but doesn't incur the overhead of creating and returning another slice.
func TakeIf[T any](slice []T, pred Predicate[T], fn func(T)) {
	for i := 0; i < len(slice); i++ {
		if pred(slice[i]) {
			fn(slice[i])
		}
	}
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

// CountBy returns the number of occurrences that satisfy the predicate.
func CountBy[T any](in []T, pred Predicate[T]) int {
	count := 0
	for i := 0; i < len(in); i++ {
		if pred(in[i]) {
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
func Map[T, R any](slice []T, mapper func(item T) R) []R {
	results := make([]R, 0)
	for i := 0; i < len(slice); i++ {
		results = append(results, mapper(slice[i]))
	}
	return results
}

// FlatMap creates a new slice mapping the values that result from applying
// the mapper function.
func FlatMap[T, R any](slice []T, mapper func(item T) []R) []R {
	results := make([]R, 0)
	for i := 0; i < len(slice); i++ {
		results = append(results, mapper(slice[i])...)
	}
	return results
}

// Accumulator is a function type use to reduce a slice to an accumulated value.
// It takes the current accumulated value and item from a slice returning the
// updated result.
type Accumulator[T, R any] func(agg R, item T) R

// Reduce reduces a slice to a value that is accumulated by iterating over each
// element in the slice.
func Reduce[T, R any](slice []T, accum Accumulator[T, R], val R) R {
	for _, item := range slice {
		val = accum(val, item)
	}
	return val
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
	random.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// Chunk accepts a slice and a size splitting the slice into chunks with a max length
// of the provided size. If the slice cannot be split evenly the last slice will contain
// all the remaining elements.
//
// Providing a size less than 1 will result in a panic.
func Chunk[T any](slice []T, size int) [][]T {
	if size < 1 {
		panic("illegal size, cannot create chunks whose size is less than 1")
	}
	chunks := make([][]T, 0)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Batch accepts a slice and a batch size returning the subset of the original slice
// according to the batch size provided.
//
// Batch can be useful when processing large volumes of data and needing to batch it
// in chunks for performance reasons.
//
// Providing a batch size less than 1 will result in a panic.
//
// DEPRECATED: Use Chunk instead. Batch wasn't the best naming choice and Chunk makes
// more sense.
func Batch[T any](slice []T, batchSize int) [][]T {
	return Chunk(slice, batchSize)
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

// Flatten accepts a slice of slices and flattens it into a new one dimensional
// slice.
func Flatten[T any](slice [][]T) []T {
	res := make([]T, 0)
	for i := range slice {
		res = append(res, slice[i]...)
	}
	return res
}

// Unique returns a new slice that doesn't contain any duplicate elements. If the
// slice contains duplicates only the first occurrence is kept.
func Unique[T comparable](in []T) []T {
	result := make([]T, 0, len(in))
	seen := make(map[T]struct{}, len(in))

	for i := 0; i < len(in); i++ {
		item := in[i]
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}

// GroupBy iterates over a slice and groups the results by the key generated from
// the grouper function.
func GroupBy[T any, U comparable](in []T, grouper func(item T) U) map[U][]T {
	result := make(map[U][]T)
	for _, item := range in {
		key := grouper(item)
		result[key] = append(result[key], item)
	}
	return result
}

// PartitionBy splits an array/slice into partitions determined by a partitioner
// function.
func PartitionBy[T any, U comparable](in []T, partitioner func(item T) U) [][]T {
	result := make([][]T, 0)
	seen := make(map[U]int)

	for _, item := range in {
		key := partitioner(item)
		idx, ok := seen[key]
		if !ok {
			idx = len(result)
			seen[key] = idx
			result = append(result, make([]T, 0))
		}
		result[idx] = append(result[idx], item)
	}
	return result
}

// Pair is a type representing a pair of values
type Pair[T, U any] struct {
	First  T
	Second U
}

// Zip accepts two arrays/slices and zip the values together returning a slice of
// Pairs. If the two arrays/slices are not of equal lengths this function will
// panic.
func Zip[T, U any](left []T, right []U) []Pair[T, U] {
	if len(left) != len(right) {
		panic("cannot zip slices of different lengths")
	}
	pairs := make([]Pair[T, U], 0, len(left))
	for idx, item := range left {
		pairs = append(pairs, Pair[T, U]{
			First:  item,
			Second: right[idx],
		})
	}
	return pairs
}

// Associate converts a slice into a map by running each element through a transformer
// which returns a key and value. If any elements generate the same key the last value
// will overwrite the current value.
func Associate[T any, K comparable, V any](in []T, transformer func(item T) (K, V)) map[K]V {
	res := make(map[K]V, len(in))
	for _, item := range in {
		k, v := transformer(item)
		res[k] = v
	}
	return res
}

// Replace iterates over the slice replacing the target item with the new item in
// place up the n times. An n value of -1 has the same effect as using ReplaceAll.
func Replace[T comparable](in []T, old T, new T, n int) {
	for i := range in {
		if in[i] == old && n != 0 {
			in[i] = new
			n--
		}
	}
}

// ReplaceAll iterates over the slice replacing all instances of old with the new
// value in place.
func ReplaceAll[T comparable](in []T, old T, new T) {
	Replace(in, old, new, -1)
}

// ReplaceIf iterates over a slice and replaces all elements that satisfy the
// predicate with the new value in place.
func ReplaceIf[T any](in []T, newVal T, pred Predicate[T]) {
	for i := range in {
		if pred(in[i]) {
			in[i] = newVal
		}
	}
}

// Concat takes an arbitrary set of slices and concatenates them in order returning
// a new slice.
func Concat[S ~[]E, E any](slices ...S) S {
	size := 0
	for _, s := range slices {
		size += len(s)
	}

	newSlice := make(S, size)
	i := 0

	for _, s := range slices {
		for j := 0; j < len(s); j, i = j+1, i+1 {
			newSlice[i] = s[j]
		}
	}

	return newSlice
}

// ForEachParallel iterates through a slice in parallel using the specified
// amount of parallelism.
func ForEachParallel[T any](slice []T, fn func(T), parallelism int) {

	if parallelism < 1 {
		panic(fmt.Errorf("parallelism less than 0 not permitted"))
	}

	wg := sync.WaitGroup{}
	wg.Add(parallelism)

	chanSize := parallelism * 4
	if chanSize > len(slice) {
		chanSize = len(slice)
	}

	queue := make(chan T, chanSize)
	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg.Done()
			for v := range queue {
				fn(v)
			}
		}()
	}

	go func() {
		for j := 0; j < len(slice); j++ {
			queue <- slice[j]
		}
		close(queue)
	}()

	wg.Wait()
}
