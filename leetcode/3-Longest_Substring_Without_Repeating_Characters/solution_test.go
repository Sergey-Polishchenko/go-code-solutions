package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/3-Longest_Substring_Without_Repeating_Characters/testdata"
)

func TestSolution(t *testing.T) {
	for _, tt := range testdata.Tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
				if got := Solution(tt.S); !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("Solution() = %v, want %v", got, tt.Want)
				}
			},
		)
	}
}
