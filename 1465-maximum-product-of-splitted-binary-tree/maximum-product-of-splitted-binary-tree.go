/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const MOD = 1_000_000_007

func getTotalSum(root *TreeNode) int64 {
    if root == nil {
        return 0
    }
    return int64(root.Val) +
        getTotalSum(root.Left) +
        getTotalSum(root.Right)
}

func maxProduct(root *TreeNode) int {
    totalSum := getTotalSum(root)
    var maxProd int64 = 0

    var dfs func(*TreeNode) int64
    dfs = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)

        subTreeSum := left + right + int64(node.Val)

        product := subTreeSum * (totalSum - subTreeSum)
        if product > maxProd {
            maxProd = product
        }

        return subTreeSum
    }

    dfs(root)
    return int(maxProd % MOD)
}