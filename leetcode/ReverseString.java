// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/

import java.lang.StringBuilder;

class ReverseString {
    public String reverseString(String s) {
        StringBuilder sb = new StringBuilder(s.length());
        for (int i = s.length()-1; i >= 0; i--) {
            sb.append(s.charAt(i));
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        ReverseString rs = new ReverseString();

        String[] tests = {
            "hello",
            "Alberto Jin",
            "A",
        };
        for (String test : tests) {
            System.out.println(test.length());
            System.out.println(test);
            System.out.println(rs.reverseString(test));
        }
    }
}
