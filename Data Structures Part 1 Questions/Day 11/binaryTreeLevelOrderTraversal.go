// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
// Memory Usage: 2.9 MB, less than 55.93% of Go online submissions for Binary Tree Level Order Traversal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxLevel = -1 

func levelOrder(root *TreeNode) [][]int {
    
    res := make([][]int, 0)
    
    levelOrderHelper(root, 0, &res)
    
    maxLevel = -1
    
    return res
    
}

func levelOrderHelper(root *TreeNode, level int, res *[][]int) {
    
    if root == nil {
        return 
    }
    
    if level > maxLevel {
        *res = append(*res, make([]int, 0))
        maxLevel += 1
    }
    
    (*res)[level] = append((*res)[level], root.Val)
    levelOrderHelper(root.Left, level+1, res)
    levelOrderHelper(root.Right, level+1, res)
    
    return
}