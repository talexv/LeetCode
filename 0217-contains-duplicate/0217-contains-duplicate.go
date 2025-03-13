// Compute: O(n)
// Memory: O(n)

func containsDuplicate(nums []int) bool {
    duplicates := make(map[int]bool, len(nums))
    for _, num := range nums {
        if duplicates[num] {
            return true
        }
        duplicates[num] = true
    }
    return false
}