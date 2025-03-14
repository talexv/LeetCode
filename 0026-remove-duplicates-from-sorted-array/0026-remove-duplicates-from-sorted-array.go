// Compute: O(n)
// Memory: O(1)

func removeDuplicates(nums []int) int {
    uniqueIndex := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[uniqueIndex] {
            uniqueIndex++
            nums[uniqueIndex] = nums[i]
        }
    }
    nums = nums[:uniqueIndex+1]
    return uniqueIndex + 1
}