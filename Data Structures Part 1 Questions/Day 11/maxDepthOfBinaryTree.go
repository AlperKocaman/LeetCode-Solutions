// Runtime: 4 ms, faster than 86.11% of Go online submissions for Maximum Depth of Binary Tree.
// Memory Usage: 4.1 MB, less than 99.14% of Go online submissions for Maximum Depth of Binary Tree.

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
    
    return max(maxDepth(root.Left), maxDepth(root.Right)) +1
}

func max(left, right int) int {
    if left > right {
        return left
    }
    
    return right
}