// 5kyu. Maximum subarray sum
package maximum_subarray_sum

// Kadane's Algorithm variation for this problem
// Dificulty O(n)
func Solution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, sum := 0, 0
	for i := range nums {
		sum += nums[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}
