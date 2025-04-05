package median_of_two_sorted_arrays_iterations

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/4-Median_of_Two_Sorted_Arrays/testdata"
)

type solution func(nums1 []int, nums2 []int) float64

var solutions = []struct {
	name string
	fn   solution
}{
	{"V1", SolutionV1},
	{"V2", SolutionV2},
	{"V3", SolutionV3},
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
					got := s.fn(tt.Nums1, tt.Nums2)
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
		nums1, nums2 := generateData(bc.size)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, nums1, nums2)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution solution, nums1 []int, nums2 []int) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(nums1, nums2)
	}
}

func generateData(size int) ([]int, []int) {
	a := make([]int, 0, size/2)
	b := make([]int, 0, size/2)
	for i := 0; i < size; i += 1 {
		if i&1 == 0 {
			a = append(a, i)
			continue
		}
		b = append(b, i)
	}
	return a, b
}
