/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHight := dfs(node.Left)
		if leftHight == -1 {
			return -1
		}

		rightHight := dfs(node.Right)
		if rightHight == -1 {
			return -1
		}

		if abs(leftHight - rightHight) > 1 {
			return -1
		}

		return 1 + max(leftHight, rightHight)
	}

	return dfs(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}