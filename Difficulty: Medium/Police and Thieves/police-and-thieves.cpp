class Solution {
  public:
    int catchThieves(vector<char> &arr, int k) {
        int n = arr.size();
        vector<int> police, thieves;
    
        // store positions
        for (int i = 0; i < n; i++) {
            if (arr[i] == 'P') police.push_back(i);
            else if (arr[i] == 'T') thieves.push_back(i);
        }
        int i = 0, j = 0, ans = 0;
        while (i < police.size() && j < thieves.size()) {
            if (abs(police[i] - thieves[j]) <= k) {
                ans++;
                i++;
                j++;
            } 
            else if (police[i] < thieves[j]) {
                i++;
            } 
            else {
                j++;
            }
        }
    
        return ans;
    }
};