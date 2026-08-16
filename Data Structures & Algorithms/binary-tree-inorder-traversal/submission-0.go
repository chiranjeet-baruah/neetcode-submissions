/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// INORDER
// LEFT - ROOT - RIGHT

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	current := root

	if root == nil {
		return []int{}
	}

	for current != nil || len(stack) > 0 {
		// Push all left children of current node to the stack
		for current != nil {
			stack = append(stack,current)
			current = current.Left
		}

		// Current is nil, pop the top element from the stack
		backIndex := len(stack) - 1
		current = stack[backIndex]
		stack = stack[:backIndex]

		// Add the node value to results
		result = append(result, current.Val)

		// Move to the right subtree
		current = current.Right
	}

	return result
}