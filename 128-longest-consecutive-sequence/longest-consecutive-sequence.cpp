class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_map<int, bool> seen;
        int ans = 0;
        for(int i = 0; i < nums.size(); i++){
            seen[nums[i]] = true;
        }
        for(auto it: seen) {
            int n = it.first;
            if(seen.find(n-1) == seen.end()){
                int j = 0;
                while(seen.find(n++) != seen.end()){
                    j++;
                }
                ans = max(ans, j);
            }
        }
        return ans;
    }
};