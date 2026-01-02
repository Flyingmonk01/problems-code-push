class Solution {
  public:
    void sort012(vector<int>& a) {
        // code here
        int n = a.size();
        int low = 0;
        int mid = 0;
        int high = n-1;
        while(mid <= high){
            if(a[mid] == 0){
                swap(a[low++], a[mid++]);
            }
            else if(a[mid] == 1){
                mid++;
            }
            else if(a[mid] == 2){
                swap(a[mid], a[high--]);
            }
        }
        
    }
};