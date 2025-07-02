package wepie

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteLastNthNode(head *ListNode, n int) *ListNode {
	left, right := head, head

	for i := 0; i < n+1; i++ {
		right = right.Next
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return head
}
