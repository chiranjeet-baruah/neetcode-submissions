/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	current := root

	if root == nil {
		return []int{}
	}

	for current != nil || len(stack) > 0{
		// Process the node and push its right child to the stack
		for current != nil {
			result = append(result,current.Val)

			if current.Right != nil {
				stack = append(stack,current.Right)
			}

			current = current.Left
		}

		// When left path ends, pop the last pushed right child to explore it
		if len(stack) > 0 {
			backIndex := len(stack)-1
			current = stack[backIndex]
			stack = stack[:backIndex]
		}
	}

	return result
}