// 4. Median of Two Sorted Arrays
package median_of_two_sorted_arrays

// TODO: Solution description
func Solution(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	low, high := 0, m
	for low <= high {
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		if i > 0 && nums1[i-1] > nums2[j] {
			high = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			low = i + 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 != 0 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
