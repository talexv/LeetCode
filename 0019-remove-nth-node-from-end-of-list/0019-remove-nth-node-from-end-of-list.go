package removenthnodefromendoflist

// Compute: O(n)
// Memory: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

//nolint:unused // solution LeetCode problem
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length int

	node := head
	for node != nil {
		length++
		node = node.Next
	}

	index := length - n
	if index == 0 {
		return head.Next
	}

	node = head
	for range index - 1 {
		node = node.Next
	}

	node.Next = node.Next.Next

	return head
}
