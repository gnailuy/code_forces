// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/745/

class PowerOfThree {
    public boolean isPowerOfThree(int n) {
        while (true) {
            if (0 == n) return false;
            else if (1 == n) return true;
            else if (0 != n%3) return false;
            else n /= 3;
        }
    }

    public static void main(String[] args) {
        PowerOfThree pt = new PowerOfThree();

        int[] tests = {
            27,
            0,
            1,
            9,
            45,
        };

        for (int test: tests) {
            System.out.println(test);
            System.out.println(pt.isPowerOfThree(test));
            System.out.println();
        }
    }
}

