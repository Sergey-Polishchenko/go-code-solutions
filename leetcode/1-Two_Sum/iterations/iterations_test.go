package two_sum_iterations

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/1-Two_Sum/testdata"
)

type solution func([]int, int) []int

var solutions = []struct {
	name string
	fn   solution
}{
	{"BruteForce", SolutionV1},
	{"Optimized", SolutionV2},
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
					got := s.fn(tt.Nums, tt.Target)
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
		nums, target := generateData(bc.size)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, nums, target)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution solution, nums []int, target int) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(nums, target)
	}
}

func generateData(size int) ([]int, int) {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i
	}
	return nums, nums[size-1] + nums[size-2]
}
