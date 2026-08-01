/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leftHeight(root *TreeNode) int {
    h := 0
    for root != nil {
        h++
        root = root.Left
    }
    return h
}

func rightHeight(root *TreeNode) int {
    h := 0
    for root != nil {
        h++
        root = root.Right
    }
    return h
}
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lh := leftHeight(root)
    rh := rightHeight(root)

    if lh == rh {
        return (1 << lh) - 1
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}