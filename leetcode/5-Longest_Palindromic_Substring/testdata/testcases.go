package testdata

var Tests = []struct {
	Name string
	S    string
	Want string
}{
	{
		Name: "single character",
		S:    "a",
		Want: "a",
	},
	{
		Name: "full odd-length palindrome",
		S:    "racecar",
		Want: "racecar",
	},
	{
		Name: "full even-length palindrome",
		S:    "abba",
		Want: "abba",
	},
	{
		Name: "multiple same characters",
		S:    "ccc",
		Want: "ccc",
	},
	{
		Name: "palindrome at start",
		S:    "abacdfgdcaba",
		Want: "aba",
	},
	{
		Name: "palindrome at end",
		S:    "abcdc",
		Want: "cdc",
	},
	{
		Name: "numbers and symbols",
		S:    "12@21#",
		Want: "12@21",
	},
	{
		Name: "multiple palindromes same length",
		S:    "abcbcb",
		Want: "bcbcb",
	},
	{
		Name: "empty string",
		S:    "",
		Want: "",
	},
	{
		Name: "whitespace test",
		S:    "a  b  a",
		Want: "a  b  a",
	},
	{
		Name: "longest between two options",
		S:    "abbaccc",
		Want: "abba",
	},
	{
		Name: "complex nested",
		S:    "abaxyzzyxba",
		Want: "xyzzyx",
	},
}
