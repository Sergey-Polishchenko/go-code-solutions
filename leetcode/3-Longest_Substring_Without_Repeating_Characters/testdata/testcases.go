package testdata

var Tests = []struct {
	Name string
	S    string
	Want int
}{
	{
		Name: "test case 1",
		S:    "abcabcbb",
		Want: 3,
	},
	{
		Name: "test case 2",
		S:    "bbbbb",
		Want: 1,
	},
	{
		Name: "test case 3",
		S:    "pwwkew",
		Want: 3,
	},
}
