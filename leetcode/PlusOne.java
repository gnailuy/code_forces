// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/

class PlusOne {
    public int[] plusOne(int[] digits) {
        int carry = 1;
        for (int i = digits.length-1; i >= 0; i--) {
            digits[i] += carry;
            carry = digits[i] / 10;
            if (digits[i] >= 10) {
                digits[i] %= 10;
            }
        }
        if (carry < 1) {
            return digits;
        } else {
            int[] result = new int[digits.length+1];
            result[0] = carry;
            for (int i = 0; i < digits.length; i++) {
                result[i+1] = digits[i];
            }
            return result;
        }
    }

    public static void main(String[] args) {
        PlusOne po= new PlusOne();

        int[][] tests = {
            {1, 2, 3},
            {4, 3, 2, 1},
            {6, 9, 9, 9},
            {9, 9, 9, 9},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            int[] result = po.plusOne(test);
            System.out.println(result.length);
            for (int i = 0; i < result.length; i++) {
                System.out.print(result[i] + " ");
            }
            System.out.println();
        }
    }
}
