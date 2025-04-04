package testdata

var Tests = []struct {
	Name   string
	Nums   []int
	Target int
	Want   []int
}{
	{
		Name:   "test case 1",
		Nums:   []int{2, 7, 11, 15},
		Target: 9,
		Want:   []int{0, 1},
	},
	{
		Name:   "test case 2",
		Nums:   []int{3, 2, 4},
		Target: 6,
		Want:   []int{1, 2},
	},
	{
		Name:   "test case 3",
		Nums:   []int{3, 3},
		Target: 6,
		Want:   []int{0, 1},
	},
	{
		Name:   "test case 4",
		Nums:   []int{1, 3, 2},
		Target: 3,
		Want:   []int{0, 2},
	},
	{
		Name:   "test no solution (impossible case)",
		Nums:   []int{1, 2},
		Target: 4,
		Want:   nil,
	},
}
