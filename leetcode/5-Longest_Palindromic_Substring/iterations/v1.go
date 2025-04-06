package longest_palindromic_substring_iterations

// Bad sulution for Longest Palindromic Substring
// Dificulty: O(n^3)
func SolutionV1(s string) string {
	var max int
	var mi, mj int
	for i := len(s); i > 0; i -= 1 {
		for j := 0; j <= i; j += 1 {
			sub := s[j:i]
			if isPolyndrom(sub) && len(sub) > max {
				max = len(sub)
				mj, mi = j, i
			}
		}
	}
	return s[mj:mi]
}

func isPolyndrom(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		if l <= len(s)/2 {
			l += 1
			r -= 1
		}
	}
	return true
}
