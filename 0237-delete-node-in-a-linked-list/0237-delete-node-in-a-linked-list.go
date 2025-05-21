package deletenodeinalinkedlist

// Compute: O(1)
// Memory: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

//nolint:unused // solution LeetCode problem
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
