package longest_palindromic_substring_iterations

// Manacher's algorithm variation for Longest Palindromic Substring
// Dificulty: O(n) (but not effective by allocations)
func SolutionV2(s string) string {
	if len(s) < 2 {
		return s
	}

	t := make([]byte, len(s)*2+3)
	t[0] = '^'
	for i := 0; i < len(s); i++ {
		t[2*i+1] = '#'
		t[2*i+2] = s[i]
	}
	t[len(t)-2] = '#'
	t[len(t)-1] = '$'

	n := len(t)
	p := make([]int, n)
	C, R := 0, 0
	maxLen, center := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*C - i
		if i < R {
			p[i] = min(R-i, p[mirror])
		}

		for t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}

		if i+p[i] > R {
			C, R = i, i+p[i]
			if p[i] > maxLen {
				maxLen = p[i]
				center = i
			}
		}
	}

	start := (center - maxLen) / 2
	return s[start : start+maxLen]
}
