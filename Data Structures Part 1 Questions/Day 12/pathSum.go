// Runtime: 4 ms, faster than 93.92% of Go online submissions for Path Sum.
// Memory Usage: 4.5 MB, less than 100.00% of Go online submissions for Path Sum.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    
    if root == nil {
        return false
    }
    
    if isLeaf(root) && targetSum == root.Val {
        
        return true
    } 
    
    return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}

func isLeaf(node *TreeNode) bool {
    if node.Left == nil && node.Right == nil {
        return true
    }
    
    return false
}