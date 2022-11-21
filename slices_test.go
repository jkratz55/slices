package slices

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  []string
		fn    func(item string) bool
	}{
		{
			name:  "Find items that start with A",
			slice: []string{"Apple", "apple", "banana", "Acorn", "Zebra"},
			want:  []string{"Apple", "Acorn"},
			fn: func(item string) bool {
				if strings.HasPrefix(item, "A") {
					return true
				}
				return false
			},
		},
		{
			name:  "Nil slice",
			slice: nil,
			want:  []string{},
			fn: func(item string) bool {
				if strings.HasPrefix(item, "A") {
					return true
				}
				return false
			},
		},
		{
			name:  "Empty slice",
			slice: []string{},
			want:  []string{},
			fn: func(item string) bool {
				if strings.HasPrefix(item, "A") {
					return true
				}
				return false
			},
		},
	}
	for _, test := range tests {
		actual := Filter[string](test.slice, test.fn)
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "Reverse Slice with Even Items",
			arg:  []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "Reverse Slice with Odd Items",
			arg:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Reverse nil slice",
			arg:  nil,
			want: nil,
		},
		{
			name: "Reverse Empty Slice",
			arg:  []int{},
			want: []int{},
		},
	}
	for _, test := range tests {
		Reverse(test.arg)
		if diff := cmp.Diff(test.arg, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestFindFirst(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		fn    func(i int) bool
		want  int
		found bool
	}{
		{
			name:  "Find First Even Number",
			slice: []int{3, 6, 8, 1, 9, 4, 1, 9},
			fn: func(i int) bool {
				if i%2 == 0 {
					return true
				}
				return false
			},
			want:  6,
			found: true,
		},
		{
			name:  "First First Number Divisible by 9",
			slice: []int{3, 6, 8, 1, 9, 4, 1, 9},
			fn: func(i int) bool {
				if i%9 == 0 {
					return true
				}
				return false
			},
			want:  9,
			found: true,
		},
		{
			name:  "Condition Not Meet",
			slice: []int{3, 6, 8, 1, 9, 4, 1, 9},
			fn: func(i int) bool {
				if i > 100 {
					return true
				}
				return false
			},
			want:  0,
			found: false,
		},
		{
			name:  "Nil Slice",
			slice: nil,
			fn: func(i int) bool {
				return false
			},
			want:  0,
			found: false,
		},
	}
	for _, test := range tests {
		actual, ok := FindFirst(test.slice, test.fn)
		if ok != test.found {
			t.Errorf("Test %s failed", test.name)
			t.Errorf("Expected found %t but got %t", test.found, ok)
		}
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestFindAll(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		fn   func(s string) bool
		want []string
	}{
		{
			name: "Find all words that start with D or d",
			arg:  []string{"David", "david", "Harley", "Davidson", "Joe", "Pizza"},
			fn: func(s string) bool {
				if strings.HasPrefix(s, "D") || strings.HasPrefix(s, "d") {
					return true
				}
				return false
			},
			want: []string{"David", "david", "Davidson"},
		},
		{
			name: "Conditions not met",
			arg:  []string{"David", "david", "Harley", "Davidson", "Joe", "Pizza"},
			fn: func(s string) bool {
				if len(s) > 20 {
					return true
				}
				return false
			},
			want: []string{},
		},
		{
			name: "Nil Slice Argument",
			arg:  nil,
			fn: func(s string) bool {
				if strings.HasPrefix(s, "D") || strings.HasPrefix(s, "d") {
					return true
				}
				return false
			},
			want: []string{},
		},
		{
			name: "Empty Slice",
			arg:  []string{},
			fn: func(s string) bool {
				if strings.HasPrefix(s, "D") || strings.HasPrefix(s, "d") {
					return true
				}
				return false
			},
			want: []string{},
		},
	}
	for _, test := range tests {
		actual := FindAll(test.arg, test.fn)
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		arg   []string
		find  string
		found bool
	}{
		{
			name:  "Value is Contained",
			arg:   []string{"I", "am", "spider", "man"},
			find:  "man",
			found: true,
		},
		{
			name:  "Value is Not Contained",
			arg:   []string{"I", "am", "spider", "man"},
			find:  "Canada",
			found: false,
		},
		{
			name:  "Nil Slice",
			arg:   nil,
			find:  "X",
			found: false,
		},
	}
	for _, test := range tests {
		found := Contains(test.arg, test.find)
		if found != test.found {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(cmp.Diff(found, test.found))
		}
	}
}
func TestMap(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
		fn   func(item int) int
	}{
		{
			name: "Map by Multiplying * 2",
			arg:  []int{1, 2, 3, 4, 5},
			want: []int{2, 4, 6, 8, 10},
			fn: func(item int) int {
				return item * 2
			},
		},
	}
	for _, test := range tests {
		actual := Map(test.arg, test.fn)
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestBatch(t *testing.T) {
	tests := []struct {
		name      string
		arg       []int
		batchSize int
		want      [][]int
	}{
		{
			name:      "Perfectly Dividable Batch",
			arg:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			batchSize: 5,
			want: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
		},
		{
			name:      "Spill over batch",
			arg:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			batchSize: 4,
			want: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15},
			},
		},
		{
			name:      "Slice Len < Batch Size",
			arg:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			batchSize: 20,
			want:      [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
		{
			name:      "Slice Len Equals Batch Size",
			arg:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			batchSize: 15,
			want:      [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
	}
	for _, test := range tests {
		actual := Batch(test.arg, test.batchSize)
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
func TestShuffle(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := make([]int, len(x))
	copy(y, x)
	Shuffle(y)
	if diff := cmp.Diff(y, x); diff == "" {
		t.Errorf("Shuffled slice is the same, this is unexpected but mathmatically possible")
		t.Errorf(diff)
	}
}
func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		arg      []int
		toRemove int
		want     []int
		removed  int
	}{
		{
			name:     "Remove Single In Middle",
			arg:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			toRemove: 5,
			want:     []int{1, 2, 3, 4, 6, 7, 8, 9},
			removed:  1,
		},
		{
			name:     "Remove Multiple",
			arg:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			toRemove: 1,
			want:     []int{2, 2, 3, 3, 4, 4, 5, 5},
			removed:  2,
		},
		{
			name:     "Remove All",
			arg:      []int{1, 1, 1, 1, 1, 1},
			toRemove: 1,
			want:     []int{},
			removed:  6,
		},
		{
			name:     "Remove Last Element",
			arg:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			toRemove: 9,
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			removed:  1,
		},
	}
	for _, test := range tests {
		actual, actualRemoved := Remove(test.arg, test.toRemove)
		if actualRemoved != test.removed {
			t.Errorf("Test Case %s failed", test.name)
			t.Errorf("Count removed doesn't match, wanted %d, got %d", test.toRemove, actualRemoved)
		}
		if diff := cmp.Diff(actual, test.want); diff != "" {
			t.Errorf("Test Case %s failed", test.name)
			t.Errorf(diff)
		}
	}
}
