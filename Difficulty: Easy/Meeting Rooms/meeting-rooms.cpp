class Solution {
  public:
    bool canAttend(vector<vector<int>> &arr) {
        // Code Here
        sort(arr.begin(), arr.end(), [](const vector<int>& a, const vector<int>& b){
            return a[0] < b[0];
        });
        
        for(int i = 1; i < arr.size(); i++){
            if(arr[i][0] < arr[i-1][1]){
                return false;
            }
        }
        
        return true;
    }
};