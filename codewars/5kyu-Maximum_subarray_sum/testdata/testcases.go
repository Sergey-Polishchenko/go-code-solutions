package testdata

var Tests = []struct {
	Name string
	Nums []int
	Want int
}{
	{
		Name: "test case 1",
		Nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		Want: 6,
	},
	{
		Name: "test case 2",
		Nums: []int{},
		Want: 0,
	},
}
