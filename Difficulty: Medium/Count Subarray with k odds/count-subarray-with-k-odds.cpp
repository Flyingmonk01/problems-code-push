class Solution {
  public:
    
    int atMost(vector<int>& arr, int k) {
        int i = 0, j = 0;
        int count = 0;
        int sum = 0;
    
        while (i < arr.size()) {
            sum += arr[i];
    
            while (sum > k) {
                sum -= arr[j];
                j++;
            }
    
            count += (i - j + 1);
            i++;
        }
    
        return count;
    }
    
    int countSubarrays(vector<int>& arr, int k) {
        // code here
        for(int i = 0; i < arr.size(); i++){
            arr[i] = arr[i]%2;
        }
        return atMost(arr, k) - atMost(arr, k-1);
    }
};