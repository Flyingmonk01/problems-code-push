class Solution {
public:
    int m, n;

    int solve(int i, int j, int neu, vector<vector<int>>& coins, vector<vector<vector<int>>>& dp){
        if(i == m-1 && j == n-1){
            if(coins[i][j] < 0 && neu > 0){
                return 0;
            }
            return coins[i][j];
        }

        if(i >= m || j >= n){
            return INT_MIN;
        }

        if(dp[i][j][neu] != INT_MIN){
            return dp[i][j][neu];
        }

        int take = coins[i][j] + max(solve(i, j+1, neu, coins, dp), solve(i+1, j, neu, coins, dp));

        int skip = INT_MIN;
        if(coins[i][j] < 0 && neu > 0){
            int rightSkip = solve(i, j + 1, neu - 1, coins, dp);
            int downSkip  = solve(i + 1, j, neu - 1, coins, dp);
            skip = max(rightSkip, downSkip);
        }
        return dp[i][j][neu] = max(skip, take);
    }

    int maximumAmount(vector<vector<int>>& coins) {
        m = coins.size();
        n = coins[0].size();

        vector<vector<vector<int>>> dp(m+1, vector<vector<int>> (n+1, vector<int>(3, INT_MIN)));

        return solve(0, 0, 2, coins, dp);
    }
};