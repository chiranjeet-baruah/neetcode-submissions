/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var previous *ListNode = nil
    var next *ListNode = nil
    current := head

    for current != nil {
        next = current.Next
        current.Next = previous
        previous = current
        current = next
    }
    return previous
}