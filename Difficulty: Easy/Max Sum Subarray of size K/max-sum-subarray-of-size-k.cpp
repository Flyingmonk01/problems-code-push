class Solution {
  public:
    int maxSubarraySum(vector<int>& arr, int k) {
        int maxSum = 0;

        // initial window
        for (int i = 0; i < k; i++) {
            maxSum += arr[i];
        }

        int windowSum = maxSum;

        for (int i = k; i < arr.size(); i++) {
            windowSum = windowSum + arr[i] - arr[i - k];
            maxSum = max(maxSum, windowSum);
        }

        return maxSum;
    }
};
