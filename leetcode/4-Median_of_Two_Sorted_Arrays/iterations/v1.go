package median_of_two_sorted_arrays_iterations

import (
	"sort"
)

// Bad solution with alocation new array M+N for Median of Two Sorted Arrays
// Deficulty: O((m+n)*log(m+n))
func SolutionV1(nums1, nums2 []int) float64 {
	mn := len(nums1) + len(nums2)
	if mn == 0 {
		return 0
	}
	nums := make([]int, 0, mn)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	sort.Slice(
		nums,
		func(i, j int) bool {
			return nums[i] < nums[j]
		},
	)
	mid := mn / 2
	if mn&1 == 0 {
		return float64(nums[mid-1]+nums[mid]) / 2
	}
	return float64(nums[mid])
}
