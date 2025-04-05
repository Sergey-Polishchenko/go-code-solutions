package longest_substring_without_repeating_characters_iterations

// Solution without strings package for Longest Substring Without Repeating Characters
// Difficulty: O(n)
func SolutionV2(s string) int {
	runePos := map[rune]int{}
	res, left := 0, 0
	for i, r := range s {
		left = max(left, runePos[r])
		runePos[r] = i + 1
		res = max(res, i-left+1)
	}
	return res
}
