package two_sum_iterations

// Optimised variation of Two Sum solution
// Difficulty: O(n)
func SolutionV2(nums []int, target int) []int {
	pairs := make(map[int]int, len(nums))
	for i, v := range nums {
		if n, ok := pairs[v]; ok && n != i {
			return []int{n, i}
		}
		pairs[target-v] = i
	}
	return nil
}
