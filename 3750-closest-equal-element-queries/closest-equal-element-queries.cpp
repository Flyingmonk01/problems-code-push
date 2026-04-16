class Solution {
public:

    vector<int> solveQueries(vector<int>& nums, vector<int>& queries) {
        unordered_map<int, vector<int>> idxMp;
        vector<int> ans;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            idxMp[nums[i]].push_back(i);
        }

         for (int q : queries) {
            vector<int>& v = idxMp[nums[q]];

            // only one occurrence
            if (v.size() == 1) {
                ans.push_back(-1);
                continue;
            }

            int pos = lower_bound(v.begin(), v.end(), q) - v.begin();
            int res = INT_MAX;

            // left neighbor
            int left = v[(pos - 1 + v.size()) % v.size()];
            int d1 = abs(q - left);
            res = min(res, min(d1, n - d1));

            // right neighbor
            int right = v[(pos + 1) % v.size()];
            int d2 = abs(q - right);
            res = min(res, min(d2, n - d2));

            ans.push_back(res);
        }
        return ans;
    }
};