// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/794/

import java.util.LinkedList;
import java.util.List;

class GenerateParentheses {
    private void parenthesisHelper(String prefix, int balance,
            int left, int right, List<String> result) {
        if (balance < 0) return;

        if (left == 0 && right == 0) {
            result.add(prefix);
        } else {
            if (left > 0) {
                parenthesisHelper(prefix + "(", balance+1, left-1, right, result);
            }
            if (right > 0) {
                parenthesisHelper(prefix + ")", balance-1, left, right-1, result);
            }
        }
    }

    public List<String> generateParenthesis(int n) {
        List<String> result = new LinkedList<String>();
        parenthesisHelper("", 0, n, n, result);
        return result;
    }

    public static void main(String[] args) {
        GenerateParentheses gp = new GenerateParentheses();

        int[] tests = {
            0, 1, 2, 3, 4, 5
        };
        for (int test : tests) {
            System.out.println(test);
            List<String> parentheses = gp.generateParenthesis(test);
            for (String s: parentheses) {
                System.out.println(s);
            }
            System.out.println();
        }
    }
}
