package directions_reduction_iterations

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/codewars/5kyu-Directions_Reduction/testdata"
)

type solution func([]string) []string

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
					got := s.fn(tt.Arr)
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
		arr := generateData(bc.size)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, arr)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution solution, arr []string) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(arr)
	}
}

func generateData(size int) []string {
	cases := [4]string{"SOUTH", "EAST", "WEST", "NORTH"}
	res := make([]string, 0, size)
	for i := range size {
		res = append(res, cases[i%4])
	}
	return res
}
