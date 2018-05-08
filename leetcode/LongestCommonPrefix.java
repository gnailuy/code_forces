// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/

import java.lang.StringBuilder;

class LongestCommonPrefix {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length <= 0 || strs[0].length() <= 0) {
            return "";
        }

        int[] indexes = new int[strs.length]; // Zeros
        while (true) {
            boolean all_equal = true;
            for (int i = 0; i < strs.length; i++) {
                if (strs[i].length() <= indexes[i] || strs[i].charAt(indexes[i]) != strs[0].charAt(indexes[i])) {
                    all_equal = false;
                    break;
                }
            }
            if (all_equal) {
                for (int i = 0; i < strs.length; i++) {
                    indexes[i]++;
                }
            } else {
                break;
            }
        }

        int len = indexes[0];
        if (len > 0) {
            StringBuilder sb = new StringBuilder(len);
            for (int i = 0; i < len; i++) {
                sb.append(strs[0].charAt(i));
            }
            return sb.toString();
        } else {
            return "";
        }
    }

    public static void main(String[] args) {
        LongestCommonPrefix lcp = new LongestCommonPrefix();

        String[][] tests = {
            {"flower", "flow", "flight"},
            {"dog", "racecar", "car"},
            {"abc", "abd", "a"},
            {},
            {"a"},
            {"ab"},
            {"", "abc"},
            {"abc", "abd", ""},
        };
        for (String[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            String result = lcp.longestCommonPrefix(test);
            System.out.println(result.length());
            System.out.println(result);
            System.out.println();
        }
    }
}
