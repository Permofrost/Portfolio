#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        if(nums.size() <2){return false;}

        std::sort(nums.begin(),nums.end());
        for (int i = 0; i < nums.size()-1; i++){
            if (nums[i] == nums[i+1]){
                return true;
            }
        }
        return false;
    }
};

int main() {
    Solution solution;
    vector<int> nums = {4, 2, 3, 1};
    bool result = solution.containsDuplicate(nums);
    cout << (result ? "true" : "false") << endl;
    return 0;
}