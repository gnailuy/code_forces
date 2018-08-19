// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/820/

class DivideTwoIntegers {
    public int divide(int dividend, int divisor) {
        if (1 == divisor) return dividend;
        else if (-1 == divisor && dividend == Integer.MIN_VALUE) return Integer.MAX_VALUE;
        else if (-1 == divisor) return -dividend;

        int sign = ((dividend < 0) ^ (divisor < 0)) ? -1 : 1;

        if (dividend > 0) dividend = -dividend;
        if (divisor > 0) divisor = -divisor;

        int count = 0;
        while (dividend <= divisor) {
            int divisorMultiple = divisor;
            int currentCount = 1;
            while (dividend - divisorMultiple <= divisorMultiple) {
                divisorMultiple <<= 1; // *= 2
                currentCount <<= 1;
            }
            dividend -= divisorMultiple;
            count += currentCount;
        }
        return sign*count;
    }

    public static void main(String[] args) {
        DivideTwoIntegers dti = new DivideTwoIntegers();

        int[][] tests = {
            {2, 2},
            {3, 3},
            {10, 3},
            {7, 3},
            {7, -3},
            {-7, -3},
            {-7, 3},
            {Integer.MIN_VALUE, 3},
            {Integer.MAX_VALUE, 3},
            {Integer.MIN_VALUE, 2},
            {Integer.MAX_VALUE, 2},
            {Integer.MIN_VALUE, 1},
            {Integer.MAX_VALUE, 1},
            {1120525167, -996222521},
            {-1718798136, 578319093},
        };

        for (int[] test: tests) {
            System.out.println(test[0] + "/" + test[1] + " = " + dti.divide(test[0], test[1]));
        }
        System.out.println();
    }
}

