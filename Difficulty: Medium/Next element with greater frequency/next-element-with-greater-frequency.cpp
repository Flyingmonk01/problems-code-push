class Solution {
  public:
    vector<int> nextFreqGreater(vector<int>& arr) {
        // code here
        int n = arr.size();
        unordered_map<int, int> freq;
        for (int x : arr) {
            freq[x]++;
        }
        
        vector<int> ans(n, -1);
        stack<int> st; 
        
        for (int i = n - 1; i >= 0; i--) {
        // Remove elements with frequency <= current
            while (!st.empty() && freq[arr[st.top()]] <= freq[arr[i]]) {
                st.pop();
            }
    
            // If stack not empty, top is the answer
            if (!st.empty()) {
                ans[i] = arr[st.top()];
            }
    
            // Push current index
            st.push(i);
        }
        
        return ans;
        
    }
};
