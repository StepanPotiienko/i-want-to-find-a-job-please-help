class Solution {
    public:
        bool isPalindrome(int x) {
            string x_str = to_string(x);

            return x_str == string(x_str.rbegin(), x_str.rend());
        }
};