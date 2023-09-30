package slices

import (
	"testing"
)

func generateDataSet(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

func benchmarkSimulator(i int) {
	// do nothing
}

func BenchmarkForEachParallel(b *testing.B) {
	b.ReportAllocs()

	tests := []struct {
		name        string
		input       []int
		fn          func(int)
		parallelism int
	}{
		{
			name:        "100000 Elements with Parallelism of 4",
			input:       generateDataSet(100000),
			fn:          benchmarkSimulator,
			parallelism: 4,
		},
		{
			name:        "100000 Elements with Parallelism of 8",
			input:       generateDataSet(100000),
			fn:          benchmarkSimulator,
			parallelism: 8,
		},
		{
			name:        "1000000 Elements with Parallelism of 4",
			input:       generateDataSet(1000000),
			fn:          benchmarkSimulator,
			parallelism: 4,
		},
		{
			name:        "1000000 Elements with Parallelism of 8",
			input:       generateDataSet(1000000),
			fn:          benchmarkSimulator,
			parallelism: 8,
		},
		{
			name:        "10000000 Elements with Parallelism of 4",
			input:       generateDataSet(10000000),
			fn:          benchmarkSimulator,
			parallelism: 4,
		},
		{
			name:        "10000000 Elements with Parallelism of 4",
			input:       generateDataSet(10000000),
			fn:          benchmarkSimulator,
			parallelism: 8,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ForEachParallel(test.input, test.fn, test.parallelism)
			}
		})
	}
}
