package lintcode

import "fmt"

func RemoveElements(head *ListNode, val int) *ListNode { //lintcode 452
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

func RemoveElements2(head *ListNode, val int) *ListNode { //lintcode 452 使用递归解决

	if head == nil {
		return nil
	}
	res := RemoveElements2(head.Next, val)
	if head.Val == val {
		return res
	} else {
		head.Next = res
		return head
	}
}

func testThis() {
	head := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				6,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							5,
							&ListNode{
								6,
								&ListNode{
									2,
									&ListNode{
										1,
										&ListNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	head = *RemoveElements2(&head, 6)
	fmt.Println(head.Val)
	fmt.Println(head.Next.Val)
	fmt.Println(head.Next.Next.Val)
	fmt.Println(head.Next.Next.Next.Val)
	fmt.Println(head.Next.Next.Next.Next.Val)
}
