// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/836/

import java.util.LinkedList;

class BasicCalculator {
    private int doCalculate(char op, int a, int b) {
        System.out.println(a + " " + op + " " + b);
        if ('+' == op) return a+b;
        else if ('-' == op) return a-b;
        else if ('*' == op) return a*b;
        else return a/b;
    }

    public int calculate(String s) {
        LinkedList<Integer> numStack = new LinkedList<Integer>();
        LinkedList<Character> opStack = new LinkedList<Character>();

        int i = 0;
        while (i < s.length()) {
            if (' ' == s.charAt(i)) {
                i++;
            } else if ('*' == s.charAt(i) || '/' == s.charAt(i)) {
                while (!opStack.isEmpty() && ('*' == opStack.peek() || '/' == opStack.peek())) {
                    char op = opStack.pop();
                    int b = numStack.pop();
                    int a = numStack.pop();
                    numStack.push(doCalculate(op, a, b));
                }

                opStack.push(s.charAt(i));
                i++;
            } else if ('+' == s.charAt(i) || '-' == s.charAt(i)) {
                while (!opStack.isEmpty()) {
                    char op = opStack.pop();
                    int b = numStack.pop();
                    int a = numStack.pop();
                    numStack.push(doCalculate(op, a, b));
                }

                opStack.push(s.charAt(i));
                i++;
            } else {
                int num = 0;
                while (i < s.length() && '0' <= s.charAt(i) && s.charAt(i) <= '9') {
                    num *= 10;
                    num += s.charAt(i) - '0';
                    i++;
                }
                numStack.push(num);
            }
        }
        while (!opStack.isEmpty()) {
            char op = opStack.pop();
            int b = numStack.pop();
            int a = numStack.pop();
            numStack.push(doCalculate(op, a, b));
        }
        return numStack.pop();
    }

    public static void main(String[] args) {
        BasicCalculator bc = new BasicCalculator();

        String[] tests = {
            "3+2*2",
            " 3/2 ",
            " 3+5 / 2 ",
            "100000000/1/2/3/4/5/6/7/8/9/10",
        };
        for (String test : tests) {
            System.out.println(test);
            System.out.println(bc.calculate(test));
            System.out.println();
        }
    }
}
