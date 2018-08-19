// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/815/

import java.util.HashSet;

class HappyNumber {
    public boolean isHappy(int n) {
        if (n == 1) return true;

        HashSet<Integer> met = new HashSet<Integer>();
        met.add(n);

        while (true) {
            int next = 0;
            while (n > 0) {
                int digit = n%10;
                n /= 10;
                next += digit*digit;
            }

            if (1 == next) {
                return true;
            } else if (met.contains(next)) {
                return false;
            } else {
                met.add(next);
                n = next;
            }
        }
    }

    public static void main(String[] args) {
        HappyNumber hn = new HappyNumber();

        for (int i = 0; i <= 100; i++) {
            System.out.println(i + ": " + hn.isHappy(i));
        }
    }
}

