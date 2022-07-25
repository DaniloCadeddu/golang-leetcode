package reversedlinkedlist

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}
