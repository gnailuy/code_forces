// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/

class ReverseIntegerNaive {
    public int reverse(int x) {
        boolean negative = false;
        if (x < 0) {
            negative = true;
            x = -x;
        }
        if (x < 0) return 0; // For -2147483648

        int r = 0;
        while (x > 0) {
            if (Integer.MAX_VALUE / 10 < r) {
                return 0;
            }
            r = r*10 + x%10;
            x /= 10;
            if (r < 0) {
                return 0;
            }
        }

        if (!negative) return r;
        else return -r;
    }

    public static void main(String[] args) {
        ReverseIntegerNaive ri = new ReverseIntegerNaive();

        int[] tests = {
            -1,
            123,
            -123,
            120,
            12000,
            2147483647,
            -2147483647,
            -2147483648,
            0,
            1534236469,
        };
        for (int test : tests) {
            System.out.println(test);
            System.out.println(ri.reverse(test));
            System.out.println();
        }
    }
}
