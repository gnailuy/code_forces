// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/818/

class MyPow {
    public double myPow(double x, int n) {
        int sign = 1;
        if (x < 0) {
            x = -x;
            sign = n%2 == 0 ? 1 : -1;
        }
        if (0 == n || 1.0 == x) {
            return sign * 1.0; // sign is 1 when n is 0
        }

        if (n < 0) {
            if (n == Integer.MIN_VALUE) return myPow(1.0/x, Integer.MAX_VALUE);
            else return myPow(1.0/x, -n);
        } else {
            double result = 1.0;
            for (int i = 0; i < n; i++) {
                result *= x;
                if (x < 1 && result == 0.0) return 0.0;
                if (x > 1 && Double.MAX_VALUE/x < result) return sign * Double.MAX_VALUE;
            }
            return sign * result;
        }
    }

    public static void main(String[] args) {
        MyPow mp = new MyPow();

        double[] xs = {
            2.00000,
            2.10000,
            2.00000,
            1.00000,
            -1.00000,
        };
        int[] ns = {
            10,
            3,
            -2,
            Integer.MIN_VALUE,
            Integer.MAX_VALUE,
        };

        for (int i = 0; i < xs.length; i++) {
            System.out.println(xs[i] + "^" + ns[i] + " = " + mp.myPow(xs[i], ns[i]));
        }
        System.out.println();
    }
}
