/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// mergeTwoLists optimized with a local pointer instead of an allocated struct.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Handle base cases early to skip loop setup if a list is empty.
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Determine the true head node to avoid a dummy struct allocation.
	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	// Track the tail using a direct local variable pointer.
	tail := head

	// Main comparison loop.
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Direct assignment for remaining nodes instead of conditional checking.
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return head
}