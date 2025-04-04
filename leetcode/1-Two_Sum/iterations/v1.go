package two_sum_iterations

// Brutforce variation of Two Sum solution
// Difficulty: O(n^2)
func SolutionV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
