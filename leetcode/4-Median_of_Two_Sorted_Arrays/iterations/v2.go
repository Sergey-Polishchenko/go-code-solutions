package median_of_two_sorted_arrays_iterations

// Solution for Median of Two Sorted Arrays
// Deficulty: O(m+n)
func SolutionV2(nums1, nums2 []int) float64 {
	nums := mergeSortedArrays(nums1, nums2)
	mn := len(nums)
	if mn == 0 {
		return 0
	}
	mid := mn / 2
	if mn&1 == 0 {
		return float64(nums[mid-1]+nums[mid]) / 2
	}
	return float64(nums[mid])
}

func mergeSortedArrays(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
			continue
		}
		result = append(result, b[j])
		j++
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
