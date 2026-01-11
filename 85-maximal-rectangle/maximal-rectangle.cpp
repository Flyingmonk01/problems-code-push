class Solution {
public:
    vector<int> getNSL(vector<int>& h){
        stack<int>st;
        vector<int> NSL(h.size());
        for(int i = 0; i < h.size(); i++){
            if(st.empty()){
                NSL[i] = -1;
            }else{
                while(!st.empty() && h[i] <= h[st.top()]){
                    st.pop();
                }
                if(st.empty()){
                    NSL[i] = -1;
                }else{
                    NSL[i] = st.top();
                }
            }
            st.push(i);
        }
        return NSL;
    }
    vector<int> getNSR(vector<int>& h){
        stack<int> st;
        vector<int> NSR(h.size());
        for(int i = h.size()-1; i >= 0; i--){
            if(st.empty()){
                NSR[i] = h.size();
            }else{
                while(!st.empty() && h[i] <= h[st.top()]){
                    st.pop();
                }
                if(st.empty()){
                    NSR[i] = h.size();
                }else{
                    NSR[i] = st.top();
                }
            }
            st.push(i);
        }
        return NSR;
    }

    int getMaxArea(vector<int>& h){
        vector<int> NSL = getNSL(h);
        vector<int> NSR = getNSR(h);
        int maxArea = 0;
        vector<int> width(h.size());
        for(int i = 0; i < h.size(); i++){
            width[i] = NSR[i] - NSL[i] - 1;
        }
        for(int i = 0; i < h.size(); i++){
            maxArea = max(maxArea, width[i]*h[i]);
        }
        return maxArea;
    }

    int maximalRectangle(vector<vector<char>>& matrix) {
        int m = matrix.size();
        int n = matrix[0].size();

        vector<int> height(n);
        for(int i = 0; i < n; i++){
            height[i] = (matrix[0][i] == '1') ? 1 : 0;
        }
        int maxArea = getMaxArea(height);

        for(int row = 1; row < m; row++){
            for(int col = 0; col < n; col++){
                if(matrix[row][col] == '0'){
                    height[col] = 0;
                }else{
                    height[col] += 1;
                }
            }
            maxArea = max(maxArea, getMaxArea(height));
        }
        return maxArea;
    }
};