#include <iostream>
#include <vector>

using namespace std;

vector<int> twoSum(vector<int> &nums, int target) {
    vector<int> indices(2);

    for (int i = 0, length = nums.size(); i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (nums[i] + nums[j] == target) {
                indices[0] = i;
                indices[1] = j;
                return indices;
            }
        }
    }

    return indices;
}

int main() {
    vector<int> nums = {2, 7, 11, 15};

    int target = 9;

    vector<int> indices = twoSum(nums, target);

    cout << (indices[0] == 0 && indices[1] == 1) << endl;
}
