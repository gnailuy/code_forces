// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/780/

class LongestPalindromicSubstring {
    public String longestPalindromicSubstring(String s) {
        int pStart = 0;
        int maxPLength = 0;

        int length;
        int start;
        for (int i = 0; i < s.length(); i++) {
            length = 1;
            start = i;
            for (int j = 1; i-j >= 0 && i+j < s.length(); j++) {
                if (s.charAt(i-j) == s.charAt(i+j)) {
                    length += 2;
                    start = i-j;
                } else {
                    break;
                }
            }
            if (length > maxPLength) {
                maxPLength = length;
                pStart = start;
            }

            length = 0;
            start = i;
            for (int j = 0; i-j >= 0 && i+j+1 < s.length(); j++) {
                if (s.charAt(i-j) == s.charAt(i+j+1)) {
                    length += 2;
                    start = i-j;
                } else {
                    break;
                }
            }
            if (length > maxPLength) {
                maxPLength = length;
                pStart = start;
            }
        }

        return s.substring(pStart, pStart+maxPLength);
    }

    public static void main(String[] args) {
        LongestPalindromicSubstring lps = new LongestPalindromicSubstring();

        String[] tests = {
            "babad",
            "cbbd",
            "a",
            "aa",
            "aaa",
            "aba",
            "abb",
            "bba",
            "abba",
            "abcdefjhijkdcba",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(lps.longestPalindromicSubstring(test));
            System.out.println();
        }
    }
}
