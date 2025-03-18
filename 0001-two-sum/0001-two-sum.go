// Compute: O(n)
// Memory: O(n)

func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i, num := range nums {
        complement := target - num
        index, found := m[complement]
        if found {
            return []int{index, i}
        }
        m[num] = i
    }
    return nil
}