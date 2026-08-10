class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        std::vector<int> dNum;
        std::unordered_set<int> nums_set(nums.begin(), nums.end());
        for(int i =1 ; i <= nums.size(); i++){
            if(nums_set.find(i) == nums_set.end()){
                dNum.push_back(i);
            }

        }
        return dNum;
    }
};