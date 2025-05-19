/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

	for i := 0; i < index-1; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next

	return head
}