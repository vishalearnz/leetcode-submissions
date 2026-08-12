class Solution {
public:
    void dfs(int i, int j, vector<vector<char>> & grid, vector<vector<bool>> & visited, int r, int c) {
        
        if(i < 0 || j < 0 || i >= r || j >= c || visited[i][j] || grid[i][j] != '1'){
            return;
        }
        visited[i][j] = true;
        // for(auto i : grid){
            // for(auto j : grid[0]){
        dfs(i-1, j, grid, visited, r, c);
        dfs(i, j+1, grid, visited, r, c);
        dfs(i+1, j, grid, visited, r, c);
        dfs(i, j-1, grid, visited, r, c);
    }
    int numIslands(vector<vector<char>>& grid) {
        int island = 0;
        int r = grid.size();
        int c = grid[0].size();
        vector<vector<bool>> visited(r, vector<bool>(c, false));
        for(int i = 0; i < r ; i++){
            for(int j = 0; j < c; j++){
                if(grid[i][j] == '1' && !visited[i][j]) {
                    dfs(i, j, grid, visited, r, c);
                    island++;
                }
            }
        }
        return island;
    }
       
        
    
};