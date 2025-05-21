package reverselinkedlist

// Compute: O(n)
// Memory: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

//nolint:unused // solution LeetCode problem
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	return prev
}
