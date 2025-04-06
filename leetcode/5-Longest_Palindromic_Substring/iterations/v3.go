package longest_palindromic_substring_iterations

// Solution for Longest Palindromic Substring
// Difficulty: O(n^2)
func SolutionV3(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		len1 := expand(s, i, i)
		len2 := expand(s, i, i+1)
		currentMax := max(len1, len2)

		if currentMax > maxLen {
			maxLen = currentMax
			start = i - (currentMax-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
