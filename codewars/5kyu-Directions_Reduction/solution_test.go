package directions_reduction

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-code-solutions/codewars/5kyu-Directions_Reduction/testdata"
)

func TestSolution(t *testing.T) {
	for _, tt := range testdata.Tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
				if got := Solution(tt.Arr); !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("Solution() = %v, want %v", got, tt.Want)
				}
			},
		)
	}
}
