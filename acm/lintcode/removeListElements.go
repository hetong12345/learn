package lintcode

func RemoveElements(head *ListNode, val int) *ListNode {//lintcode 452
	// write your code here

	for head != nil && head.Val == val {
		next := head.Next
		head.Next = nil
		head = next
	}
	if head == nil {
		return head
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			del := prev.Next
			prev.Next = del.Next
			del.Next = nil
		} else {
			prev = prev.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
