class Solution {
  public:
    vector<int> maxOfSubarrays(vector<int>& arr, int k) {
        // code here
        vector<int> ans;
        int maxi = INT_MIN;
        
        for(int i = 0; i < k; i++){
            maxi = max(arr[i], maxi);
        }
        ans.push_back(maxi);
        for (int i = k; i < arr.size(); i++) {

        // if outgoing element was max, recompute
            if (arr[i - k] == maxi) {
                maxi = arr[i - k + 1];
                for (int j = i - k + 1; j <= i; j++) {
                    maxi = max(maxi, arr[j]);
                }
            } else {
                // otherwise, just compare with incoming element
                maxi = max(maxi, arr[i]);
            }
    
            ans.push_back(maxi);
        }
        return ans;
    }
};