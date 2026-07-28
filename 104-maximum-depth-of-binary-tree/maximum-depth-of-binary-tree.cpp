/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }
        queue<pair<TreeNode*, int>> q;
        q.push({root, 1});

        int ans = -1;

        while(!q.empty()) {
            auto front = q.front();
            q.pop();
            ans = max(ans, front.second);
            if(front.first->left != NULL) {
                q.push({front.first->left, 1+front.second});
            }
            if(front.first->right != NULL) {
                q.push({front.first->right, 1+front.second});
            }
        }
        return ans;
    }
};