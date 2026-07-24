class Solution {
public:
    int isValidOperators(string& ch) {
        return ch == "+" || ch == "-" || ch == "*" || ch == "/";
    }
    int evalRPN(vector<string>& tokens) {
        stack<int> st;
        for(int i = 0; i < tokens.size(); i++) {
            auto curr = tokens[i];
            if(isValidOperators(curr)) {
                int b = st.top(); st.pop();
                int a = st.top(); st.pop();

                if(curr == "+") {
                    st.push(a+b);
                }else if(curr == "-") {
                    st.push(a-b);
                }else if(curr == "*"){
                    st.push(a*b);
                }else if(curr == "/") {
                    st.push(a/b);
                }
            }else{
                st.push(stoi(curr));
            }
        }
        return st.top();
    }
};