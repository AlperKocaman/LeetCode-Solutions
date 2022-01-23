// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Postorder Traversal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
     /* 
     * in postorder traversal,
     * First, left subtree should be visited
     * Then right subtree should be visited,
     * Finally root should be visited.
    */
    
    res := []int{}
    
    if root != nil {
        res =  postorderTraversal(root.Left)
        res = append(res, postorderTraversal(root.Right)...)
        res = append(res, root.Val)
    }
    
    return res
}