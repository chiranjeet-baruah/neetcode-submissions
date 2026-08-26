/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    
    // helper function calculates max depth and updates maxDiameter
    var maxDepth func(*TreeNode) int
    maxDepth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        // Recursively find the depth of left and right subtrees
        leftDepth := maxDepth(node.Left)
        rightDepth := maxDepth(node.Right)
        
        // Current path through this node has length (leftDepth + rightDepth)
        if leftDepth + rightDepth > maxDiameter {
            maxDiameter = leftDepth + rightDepth
        }
        
        // Return the max depth from this node down to its deepest leaf
        if leftDepth > rightDepth {
            return 1 + leftDepth
        }
        return 1 + rightDepth
    }
    
    maxDepth(root)
    return maxDiameter
}