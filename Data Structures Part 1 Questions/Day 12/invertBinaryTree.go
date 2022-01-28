// Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Invert Binary Tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    
    if root == nil {
        return nil 
    }
    
    res := &TreeNode{Val: root.Val}
    
    invertTreeHelper(root, res)
    
    return res
}

func invertTreeHelper(root, res *TreeNode) {
    
    if root == nil {
        return
    }
    
    if root.Left != nil {
        right := &TreeNode{Val: root.Left.Val}
        res.Right = right
        invertTreeHelper(root.Left, right)
    }
    
    if root.Right != nil {
        left := &TreeNode{Val: root.Right.Val}
        res.Left = left
        invertTreeHelper(root.Right, left)
    }
    
}