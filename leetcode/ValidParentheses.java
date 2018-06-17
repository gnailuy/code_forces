// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/721/

class ValidParentheses {
    public boolean isValid(String s) {
        int n = s.length();
        if (n%2 != 0) return false;

        char[] stack = new char[n];
        int p = 0;
        for (int i = 0; i < n; i++) {
            char c = s.charAt(i);
            if ('(' == c || '[' == c || '{' == c) {
                stack[p++] = c;
            } else if (')' == c) {
                if (p > 0 && '(' == stack[p-1]) {
                    p--;
                } else {
                    return false;
                }
            } else if (']' == c) {
                if (p > 0 && '[' == stack[p-1]) {
                    p--;
                } else {
                    return false;
                }
            } else if ('}' == c) {
                if (p > 0 && '{' == stack[p-1]) {
                    p--;
                } else {
                    return false;
                }
            }
        }

        return 0 == p;
    }

    public static void main(String[] args) {
        ValidParentheses vp = new ValidParentheses();

        String[] tests = {
            "()",
            "()[]{}",
            "(]",
            "([)]",
            "{[]}",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(vp.isValid(test));
            System.out.println();
        }
    }
}
