// https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/819/

class MySqrt {
    public int mySqrt(int n) {
        if (n == 0) return 0;
        else if (n < 4) return 1;

        int i = 2;
        while (true) {
            int prod = i*i;
            if (prod == n) return i;
            else if (prod > n || prod < 0) return i-1;
            else i++;
        }
    }

    public int newtonSqrt(int n) {
        double nd = n;
        double k = 1;
        while (true) {
            k = (k+nd/k)/2;
            if (k*k-nd < 0.000001) {
                return (int)k;
            }
        }
    }

    private double abs(double d) {
        return d > 0 ? d : -d;
    }

    public static void main(String[] args) {
        MySqrt ms = new MySqrt();

        for (int i = 1; i <= 100; i++) {
            System.out.println(i + ": " + ms.mySqrt(i));
        }

        int[] tests = {
            2147395599,
            2147483647,
        };
        for (int i: tests) {
            System.out.println(i + ": " + ms.mySqrt(i));
        }
        System.out.println();

        for (int i = 1; i <= 100; i++) {
            System.out.println(i + ": " + ms.newtonSqrt(i));
        }
        for (int i: tests) {
            System.out.println(i + ": " + ms.newtonSqrt(i));
        }
        System.out.println();
    }
}

