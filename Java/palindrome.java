public class palindrome {
    public boolean isPalindrome(String s) {
        s = s.toLowerCase().replaceAll("[^a-z0-9]", "");
        int len = s.length();
        if(len<2){
            return true;
        }
        for(int i = 0; i < len/2; i++){
            if(s.charAt(i) != s.charAt(len-i-1)){
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        palindrome solution = new palindrome();
        String[] testCases = {
            "A man, a plan, a canal: Panama",
            "race car",
            "not a palindrome",
            // add more test cases here
        };
        for (String testCase : testCases) {
            boolean result = solution.isPalindrome(testCase);
            System.out.println("The string \"" + testCase + "\" is " + (result ? "" : "not ") + "a palindrome.");
        }
    }
}