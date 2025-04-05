package longest_substring_without_repeating_characters_iterations

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/3-Longest_Substring_Without_Repeating_Characters/testdata"
)

type solution func(string) int

var solutions = []struct {
	name string
	fn   solution
}{
	{"V1", SolutionV1},
	{"V2", SolutionV2},
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
					got := s.fn(tt.S)
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
		st := generateData(bc.size)

		b.Run(
			bc.name,
			func(b *testing.B) {
				for _, s := range solutions {
					b.Run(
						s.name,
						func(b *testing.B) {
							benchSolution(b, s.fn, st)
						},
					)
				}
			},
		)
	}
}

func benchSolution(b *testing.B, solution solution, s string) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		solution(s)
	}
}

func generateData(size int) string {
	return strings.Repeat("s", size-1) + "b"
}
