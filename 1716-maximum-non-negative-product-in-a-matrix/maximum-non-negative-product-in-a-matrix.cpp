class Solution {
public:
    int maxProductPath(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();

        vector<vector<long long>> minP(m, vector<long long>(n));
        vector<vector<long long>> maxP(m, vector<long long>(n));

        minP[0][0] = maxP[0][0] = grid[0][0];

        for(int j = 1; j < n; j++){
            minP[0][j] = minP[0][j-1] * grid[0][j];
            maxP[0][j] = maxP[0][j-1] * grid[0][j];
        }
        for(int i = 1; i < m; i++){
            minP[i][0] = minP[i-1][0] * grid[i][0];
            maxP[i][0] = maxP[i-1][0] * grid[i][0];
        }

        for(int i = 1; i < m; i++){
            for(int j = 1; j < n; j++){
                long long a = minP[i-1][j] * grid[i][j];
                long long b = maxP[i-1][j] * grid[i][j];
                long long c = minP[i][j-1] * grid[i][j];
                long long d = maxP[i][j-1] * grid[i][j];

                maxP[i][j] = max({a, b, c, d});
                minP[i][j] = min({a, b, c, d});
            }
        }
        long long res = maxP[m-1][n-1];

        if(res < 0) return -1;

        return res % 1000000007;
    }
};