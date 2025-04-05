package testdata

var Tests = []struct {
	Name  string
	Nums1 []int
	Nums2 []int
	Want  float64
}{
	{
		Name:  "Example 1 (odd length)",
		Nums1: []int{1, 3},
		Nums2: []int{2},
		Want:  2.0,
	},
	{
		Name:  "Example 2 (even length)",
		Nums1: []int{1, 2},
		Nums2: []int{3, 4},
		Want:  2.5,
	},
	{
		Name:  "One empty array (non-empty is odd)",
		Nums1: []int{},
		Nums2: []int{5},
		Want:  5.0,
	},
	{
		Name:  "One empty array (non-empty is even)",
		Nums1: []int{2, 4},
		Nums2: []int{},
		Want:  3.0,
	},
	{
		Name:  "All elements in nums1 < nums2 (odd total)",
		Nums1: []int{1, 2},
		Nums2: []int{3, 4, 5},
		Want:  3.0,
	},
	{
		Name:  "All elements in nums1 > nums2 (even total)",
		Nums1: []int{5, 6},
		Nums2: []int{1, 2},
		Want:  3.5,
	},
	{
		Name:  "Overlapping arrays (even)",
		Nums1: []int{2, 4, 6},
		Nums2: []int{1, 3, 5},
		Want:  3.5,
	},
	{
		Name:  "Single-element arrays (even)",
		Nums1: []int{1},
		Nums2: []int{2},
		Want:  1.5,
	},
	{
		Name:  "Mixed lengths (median in gap)",
		Nums1: []int{1, 3, 5, 7},
		Nums2: []int{2, 4},
		Want:  3.5,
	},
	{
		Name:  "Duplicate values",
		Nums1: []int{1, 1, 1},
		Nums2: []int{2, 2, 2},
		Want:  1.5,
	},
	{
		Name:  "Large value difference",
		Nums1: []int{100},
		Nums2: []int{200, 300},
		Want:  200.0,
	},
	{
		Name:  "Zero values",
		Nums1: []int{0, 0},
		Nums2: []int{0, 0},
		Want:  0.0,
	},
}
