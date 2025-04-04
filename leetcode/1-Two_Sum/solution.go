// 1. Two Sum
package two_sum

// Is v2 optimised solution with hash table usage.
// Difficulty: O(n)
func Solution(nums []int, target int) []int {
	pairs := make(map[int]int, len(nums))
	for i, v := range nums {
		if n, ok := pairs[v]; ok && n != i {
			return []int{n, i}
		}
		pairs[target-v] = i
	}
	return nil
}
