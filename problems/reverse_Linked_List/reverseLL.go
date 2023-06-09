package reverse_Linked_List

import . "sample/utils/linkedlist"

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
