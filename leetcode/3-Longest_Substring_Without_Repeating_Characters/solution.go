// 3. Longest Substring Without Repeating Characters
package longest_substring_without_repeating_characters

import (
	"strings"
)

// Solution for Longest Substring Without Repeating Characters
// Difficulty: O(n)
func Solution(s string) int {
	var sub string
	var maxLen int
	for _, r := range s {
		if index := strings.IndexRune(sub, r); index != -1 {
			sub = sub[index+1:]
		}
		sub += string(r)
		maxLen = max(maxLen, len(sub))
	}
	return maxLen
}
