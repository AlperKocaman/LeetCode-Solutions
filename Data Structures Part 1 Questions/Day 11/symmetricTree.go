// Runtime: 3 ms, faster than 43.05% of Go online submissions for Symmetric Tree.
// Memory Usage: 2.9 MB, less than 67.35% of Go online submissions for Symmetric Tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
    // tree will not be given empty since available nodes in tree is at least 1  
 
    return isSymmetricHelper(root.Left, root.Right)
}    

func isSymmetricHelper(left, right *TreeNode) bool {
    
    if left == nil && right == nil{
        return true
    } else if left == nil || right == nil{
        return false
    }
    
    if left.Val != right.Val {
        return false
    }
    
    return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}