class Solution {
  public:
    bool checkRedundancy(string &s) {
        // code here
        stack<char> st;
        for(int i = 0; i < s.size(); i++){
            if(s[i] != ')'){
                st.push(s[i]);
            }else{
                bool hasOperator = false;
                
                while (!st.empty() && st.top() != '(') {
                    char top = st.top();
                    if (top == '+' || top == '-' || top == '*' || top == '/'){
                        hasOperator = true;
                    }
                    st.pop();
                }
                
                if(!st.empty()){
                    st.pop();
                }
                
                if(!hasOperator){
                    return true;
                }
            }
        }
        return false;
    }
};
