#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

class Solution {
public:
    bool isAnagram(string s, string t) {
        
    }
};

int main() {
    Solution solution;
    string s = "anagram";
    string t = "nagaram";
    bool result = solution.isAnagram(s, t);
    cout << (result ? "true" : "false") << endl;
    return 0;
}