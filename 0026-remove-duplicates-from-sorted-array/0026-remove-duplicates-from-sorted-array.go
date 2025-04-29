package removeduplicatesfromsortedarray

// Compute: O(n)
// Memory: O(1)

//nolint:unused // solution LeetCode problem
func removeDuplicates(nums []int) int {
	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	return uniqueIndex + 1
}
