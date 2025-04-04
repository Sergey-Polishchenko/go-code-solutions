package maximum_subarray_sum_iterations

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/codewars/5kyu-Maximum_subarray_sum/testdata"
)

var solutions = []struct {
	name string
	fn   func([]int) int
}{
	{"Bruteforce", SolutionV1},
	{"Kadane's Algorithm", SolutionV2},
}

var benchCases = []struct {
	size int
	name string
}{
	{10, "Small"},
	{1000, "Medium"},
	{100_000, "Large"},
}

func TestSolutions(t *testing.T) {
	for _, s := range solutions {
		t.Run(
			s.name,
			func(t *testing.T) {
				for _, tt := range testdata.Tests {
					got := s.fn(tt.Nums)
					if !reflect.DeepEqual(got, tt.Want) {
						t.Errorf("%s failed for %s: got %v, want %v",
							s.name, tt.Name, got, tt.Want)
					}
				}
			},
		)
	}
}

func BenchmarkSolutions(b *testing.B) {
	for _, bc := range benchCases {
		nums := generateData(bc.size)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, nums)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution func([]int) int, nums []int) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(nums)
	}
}

func generateData(size int) []int {
	nums := make([]int, size)
	sign := -1
	for i := range nums {
		if i == size-3 {
			sign = 1
		}
		nums[i] = sign
	}
	return nums
}
