/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	// Base case
    if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	// Two pointers fast and slow
	// If they meet at any point, then cycle exists
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next

		// If they meet: there is a cycle
		if slow == fast {
			return true
		}
	}

	// If fast reaches end, then there is no cycle
	return false

}
