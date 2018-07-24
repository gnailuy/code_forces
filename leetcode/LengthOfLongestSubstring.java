// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/779/

import java.util.HashMap;

class LengthOfLongestSubstring {
    public int lengthOfLongestSubstring(String s) {
        int maxLength = 0;
        int currentLength = 0;
        HashMap<Character, Integer> charToIndex = new HashMap<Character, Integer>();
        for (int i = 0; i < s.length(); i++) {
            if (!charToIndex.containsKey(s.charAt(i)) || charToIndex.get(s.charAt(i)) < i-currentLength) {
                currentLength += 1;
            } else {
                if (currentLength > maxLength) {
                    maxLength = currentLength;
                }
                int idx = charToIndex.get(s.charAt(i));
                currentLength = i - idx;
            }
            charToIndex.put(s.charAt(i), i);
        }
        if (currentLength > maxLength) {
            maxLength = currentLength;
        }
        return maxLength;
    }

    public static void main(String[] args) {
        LengthOfLongestSubstring lls = new LengthOfLongestSubstring();

        String[] tests = {
            "abcabcbb",
            "bbbbb",
            "pwwkew",
            "abba",
            "dvdf",
            "au",
            "c",
            "",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(lls.lengthOfLongestSubstring(test));
            System.out.println();
        }
    }
}
