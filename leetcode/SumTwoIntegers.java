// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/822/

class SumTwoIntegers {
    public int getSum(int a, int b) {
        if (a > b) {
            a ^= b;
            b ^= a;
            a ^= b;
        }
        if (Integer.MIN_VALUE == a && b < 0) return Integer.MIN_VALUE;
        if (Integer.MAX_VALUE == b && a > 0) return Integer.MAX_VALUE;

        int result = 0;
        int carry = 0;
        int mod = 1;
        for (int i = 0; i < 32; i++) {
            int da = a&mod;
            int db = b&mod;
            int d;

            if ((da&db&carry) == 1) { // Three 1
                d = 1;
                carry = 1;
            } else if ((da&db) == 1 || (da&carry) == 1 || (db&carry) == 1) { // Two 1
                d = 0;
                carry = 1;
            } else if (da == 1 || db == 1 || carry == 1) { // One 1
                d = 1;
                carry = 0;
            } else {
                d = 0;
                carry = 0;
            }
            result |= (d << i);

            a >>= 1;
            b >>= 1;
        }
        return result;
    }

    public static void main(String[] args) {
        SumTwoIntegers sti = new SumTwoIntegers();

        int[][] tests = {
            {1, 2},
            {-1, 2},
            {1, -2},
            {-1, -2},
            {Integer.MIN_VALUE, -1},
            {Integer.MAX_VALUE, 1},
        };

        for (int[] test: tests) {
            System.out.println(test[0] + "+" + test[1] + " = " + sti.getSum(test[0], test[1]));
        }
        System.out.println();
    }
}

