// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/744/

import java.util.ArrayList;
import java.util.List;
import java.lang.Math;

class CountPrimes {
    public int countPrimes(int n) {
        List<Boolean> pl = new ArrayList<Boolean>(n+1);
        for (int i = 0; i <= n; i++) {
            pl.add(false);
        }
        int upper = (int)Math.sqrt(n) + 1;
        for (int i = 2; i <= upper; i++) {
            for (int j = i; j <= n; j++) {
                int k = i*j;
                if (k <= n) {
                    pl.set(k, true);
                } else {
                    break;
                }
            }
        }
        int cnt = 0;
        for (int i = 2; i < n; i++) {
            if (!pl.get(i)) cnt++;
        }
        return cnt;
    }

    public static void main(String[] args) {
        CountPrimes cp = new CountPrimes();

        int[] tests = {
            1,
            2,
            10,
            11,
            20,
            100,
            1000,
            1500000,
        };

        for (int test: tests) {
            System.out.println(test);
            System.out.println(cp.countPrimes(test));
            System.out.println();
        }
    }
}

