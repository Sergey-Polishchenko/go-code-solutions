package maximum_subarray_sum

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/codewars/5kyu-Maximum_subarray_sum/testdata"
)

func TestSolution(t *testing.T) {
	for _, tt := range testdata.Tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
				if got := Solution(tt.Nums); !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("Solution() = %v, want %v", got, tt.Want)
				}
			},
		)
	}
}
