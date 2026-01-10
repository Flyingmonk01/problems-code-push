class Solution {
  public:
    long long helper(string& s, int k){
        vector<int> freq(26, 0);
        int cnt = 0, l = 0, r = 0;
        long long ans = 0;
        
        int n = s.size();
        
        while(r < n){
            freq[s[r]-'a']++;
            if(freq[s[r]-'a'] == 1) cnt++;
            
            while(cnt > k){
                freq[s[l]-'a']--;
                if(freq[s[l]-'a'] <= 0){
                    cnt--;
                }
                l++;
            }
            ans += r-l+1;
            r++;
        }
        return ans;
    }
    
    int countSubstr(string& s, int k){
        return helper(s, k) - helper(s, k-1);
    }
};