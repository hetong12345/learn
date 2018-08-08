package main

import (
	"./lintcode"
	"fmt"
)

type student struct {
	name  string
	score int
}


func main() {

	head:=lintcode.ListNode{
		1,
		&lintcode.ListNode{
			2,
			&lintcode.ListNode{
				6,
				&lintcode.ListNode{
					3,
					&lintcode.ListNode{
						4,
						&lintcode.ListNode{
							5,
							&lintcode.ListNode{
								6,
								&lintcode.ListNode{
									2,
									&lintcode.ListNode{
										1,
										&lintcode.ListNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	head=*lintcode.RemoveElements(&head,6)
	fmt.Println(head.Val)
	fmt.Println(head.Next.Val)
	fmt.Println(head.Next.Next.Val)
	fmt.Println(head.Next.Next.Next.Val)
	fmt.Println(head.Next.Next.Next.Next.Val)
}

