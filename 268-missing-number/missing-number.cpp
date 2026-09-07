class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int arr_sum = 0;
        int n = nums.size();
        int expected_sum = n * (n+1)/2;
        for (int num :  nums) {
            arr_sum += num;
        }
        return expected_sum - arr_sum;
        
    }
};