package median_of_two_sorted_arrays

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/leetcode/4-Median_of_Two_Sorted_Arrays/testdata"
)

func TestSolution(t *testing.T) {
	for _, tt := range testdata.Tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
				if got := Solution(tt.Nums1, tt.Nums2); !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("Solution() = %v, want %v", got, tt.Want)
				}
			},
		)
	}
}
