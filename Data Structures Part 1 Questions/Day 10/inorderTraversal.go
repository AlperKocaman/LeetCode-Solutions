// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2.3 MB, less than 8.87% of Go online submissions for Binary Tree Inorder Traversal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    /* 
     * in inorder traversal,
     * First, left subtree should be visited
     * Then root should be visited,
     * Finally right subtree should be visited.
    */
    
    res := []int{}
    
    if root != nil {
        res =  inorderTraversal(root.Left)
        res = append(res, root.Val)
        res = append(res, inorderTraversal(root.Right)...)
    }
    
    return res
}