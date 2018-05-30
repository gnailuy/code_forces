// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/569/

class ClimbingStairs {
    public int climbStairs(int n) {
        if (1 == n) return 1;
        else if (2 == n) return 2;

        int[] cs = new int[n+1];
        cs[1] = 1;
        cs[2] = 2;

        for (int i = 3; i <= n; i++) {
            cs[i] = cs[i-1] + cs[i-2];
        }

        return cs[n];
    }

    public static void main(String[] args) {
        ClimbingStairs cs = new ClimbingStairs();

        int[] tests = {
            1,
            2,
            3,
            4,
            5,
            5,
            7,
            44,
        };
        for (int test : tests) {
            System.out.println(test);
            System.out.println(cs.climbStairs(test));
            System.out.println();
        }
    }
}
