// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
// Memory Usage: 2.4 MB, less than 11.22% of Go online submissions for Binary Tree Preorder Traversal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    /* 
     * in preorder traversal,
     * First, root should be visited
     * Then left subtree should be visited,
     * Finally right subtree should be visited.
    */
    
    if root != nil { 
        res := []int{root.Val}
        res = append(res, preorderTraversal(root.Left)...)
        res = append(res, preorderTraversal(root.Right)...)
        return res
    }
    
    return []int{}
}