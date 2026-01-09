/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode) (int, *TreeNode) {
    if node == nil {
        return 0, node
    }

    left_depth, left_node := dfs(node.Left)
    right_depth, right_node := dfs(node.Right)

    if left_depth > right_depth {
        return left_depth+1, left_node
    }

    if right_depth > left_depth {
        return right_depth+1, right_node
    }

    return left_depth+1, node
}
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    _, ans := dfs(root)
    return ans
}