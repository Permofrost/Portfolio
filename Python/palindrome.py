class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = ''.join(c for c in s.lower() if c.isalnum())
        length = len(s)

        if length < 2:
            return True
        
        for i in range(0, length//2):
            if s[i] != s[length-i-1]:
                return False
            
        return True




# Create an instance of the Solution class
solution = Solution()

# Test String
test_cases = [
    "A man, a plan, a canal: Panama",
    "race car",
    "not a palindrome",
    # add more test cases here
]

for test_case in test_cases:
    result = solution.isPalindrome(test_case)
    print(f'The string "{test_case}" is {"a palindrome" if result else "not a palindrome"}')
