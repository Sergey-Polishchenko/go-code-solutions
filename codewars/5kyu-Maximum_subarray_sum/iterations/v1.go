package maximum_subarray_sum_iterations

// Bruteforce variation of maximum subarray sum
// Difficulty: O(n^2) (too bad)
func SolutionV1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var res int
	for i := 0; i < l; i += 1 {
		var sum int
		for j := i; j < l; j += 1 {
			sum += nums[j]
			res = max(res, sum)
		}
	}
	return res
}
