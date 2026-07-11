/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
    
    // Define the recursive DFS function
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        // Calculate the height of left and right subtrees
        leftHeight := dfs(node.Left)
        rightHeight := dfs(node.Right)
        
        // Update the maximum diameter found so far
        currentDiameter := leftHeight + rightHeight
        if currentDiameter > maxDiameter {
            maxDiameter = currentDiameter
        }
        
        // Return the height of the current node's subtree
        return 1 + max(leftHeight, rightHeight)
    }
    
    dfs(root)
    return maxDiameter
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
