package longest_palindromic_substring_iterations

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/5-Longest_Palindromic_Substring/testdata"
)

type solution func(string) string

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
	pl := 4
	if size-pl < 0 {
		return ""
	}
	prefix := strings.Repeat("abc", (size-pl)/3)
	if remainder := (size - pl) % 3; remainder > 0 {
		prefix += "abc"[:remainder]
	}

	return prefix + strings.Repeat("x", pl)
}
