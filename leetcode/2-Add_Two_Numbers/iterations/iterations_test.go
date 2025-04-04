package add_two_numbers_iterations

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/2-Add_Two_Numbers/testdata"
)

type solution func(*testdata.ListNode, *testdata.ListNode) *testdata.ListNode

var solutions = []struct {
	name string
	fn   solution
}{
	{"V1", SolutionV1},
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
					got := s.fn(tt.L1, tt.L2)
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
		l1 := testdata.NewListNode(generateData(bc.size)...)
		l2 := testdata.NewListNode(1, 2, 3)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, l1, l2)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution solution, l1 *testdata.ListNode, l2 *testdata.ListNode) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(l1, l2)
	}
}

func generateData(size int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i % 10
	}
	return nums
}
