class Solution {
  public:
    int countAtMostK(vector<int> &arr, int k) {
        // code here
        unordered_map<int, int> mp;
        int i = 0, j = 0;
        int n = arr.size();
        int cnt = 0;
        while(i < n) {
            mp[arr[i]]++;
            
            while(mp.size() > k){
                mp[arr[j]]--;
                if (mp[arr[j]] == 0) {
                    mp.erase(arr[j]);
                }
                j++;
                
            }
            cnt += (i-j+1);
            i++;
        }
        return cnt;
    }
};