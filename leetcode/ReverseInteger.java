// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/

import java.lang.StringBuilder;

class ReverseInteger {
    public int reverse(int x) {
        boolean negative = false;
        String s = Integer.toString(x);
        if (s.startsWith("-")) {
            s = s.replaceAll("-", "");
            negative = true;
        }
        StringBuilder sb = new StringBuilder(s.length());
        for (int i = s.length()-1; i >= 0; i--) {
            sb.append(s.charAt(i));
        }
        String sr = sb.toString();
        sr = sr.replaceAll("^0*", "");
        if (sr.length() <= 0) {
            return 0;
        } else if (!negative && sr.length() == 10 && sr.compareTo("2147483647") > 0) {
            return 0;
        } else if (negative && sr.length() == 10 && sr.compareTo("2147483648") > 0) {
            return 0;
        } else if (!negative) {
            return Integer.parseInt(sr);
        } else {
            return -Integer.parseInt(sr);
        }
    }

    public static void main(String[] args) {
        ReverseInteger ri = new ReverseInteger();

        int[] tests = {
            -1,
            123,
            -123,
            120,
            12000,
            2147483647,
            -2147483648,
            0,
        };
        for (int test : tests) {
            System.out.println(test);
            System.out.println(ri.reverse(test));
            System.out.println();
        }
    }
}
