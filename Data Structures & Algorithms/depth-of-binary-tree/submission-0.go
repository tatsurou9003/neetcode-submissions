/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightDepth := maxDepth(root.Right) 
	leftDepth := maxDepth(root.Left)

    if rightDepth > leftDepth {
		return rightDepth + 1
	}
	return leftDepth + 1
}
