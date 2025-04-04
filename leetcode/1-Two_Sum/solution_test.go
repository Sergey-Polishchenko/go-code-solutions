package two_sum

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/1-Two_Sum/testdata"
)

func TestSolution(t *testing.T) {
	for _, tt := range testdata.Tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
				if got := Solution(tt.Nums, tt.Target); !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("Solution() = %v, want %v", got, tt.Want)
				}
			},
		)
	}
}
