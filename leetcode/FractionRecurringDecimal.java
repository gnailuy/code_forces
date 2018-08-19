// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/821/

import java.util.HashMap;

class FractionRecurringDecimal {
    public String fractionToDecimal(int numerator, int denominator) {
        if (numerator == Integer.MIN_VALUE) {
            if (denominator == 1) {
                return String.valueOf(numerator);
            } else if (denominator == -1) {
                return String.valueOf(Integer.MIN_VALUE).substring(1);
            }
        }

        int sign = ((numerator < 0) ^ (denominator < 0)) ? -1 : 1;
        if (numerator > 0) numerator = -numerator;
        if (denominator > 0) denominator = -denominator;

        int whole;
        if (numerator < denominator) {
            whole = numerator/denominator;
            numerator %= denominator;
        } else {
            whole = 0;
        }
        if (numerator == 0) {
            if (whole != 0) {
                return (sign < 0 ? "-" : "") + whole;
            } else {
                return "0";
            }
        }

        HashMap<Long, Integer> numIdx = new HashMap<Long, Integer>();
        StringBuilder sb = new StringBuilder();
        int idx = 0;
        int repeatStart = -1;
        long numeratorLong = numerator;
        while (true) {
            numeratorLong *= 10;
            if (numIdx.containsKey(numeratorLong)) {
                repeatStart = numIdx.get(numeratorLong);
                break;
            } else {
                numIdx.put(numeratorLong, idx++);
                sb.append(String.valueOf(numeratorLong/denominator));
                numeratorLong %= denominator;
                if (numeratorLong == 0) break;
            }
        }

        if (repeatStart >= 0) {
            sb.insert(repeatStart, '(');
            sb.append(')');
        }
        return (sign < 0 ? "-" : "") + whole + "." + sb.toString();
    }

    public static void main(String[] args) {
        FractionRecurringDecimal frd = new FractionRecurringDecimal();

        int[][] tests = {
            {1, 2},
            {2, 1},
            {2, 3},
            {1, 3},
            {-50, 8},
            {7, -12},
            {-1, -2147483648},
        };

        for (int[] test: tests) {
            System.out.println(test[0] + "/" + test[1] + " = " + frd.fractionToDecimal(test[0], test[1]));
        }
        System.out.println();
    }
}

