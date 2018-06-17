// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/722/

import java.util.HashSet;

class MissingNumber {
    public int missingNumber(int[] nums) {
        HashSet<Integer> partener = new HashSet<Integer>();
        int n = nums.length;
        int middle = -1;
        if (0 == n%2) {
            middle = n/2;
        }
        for (int i = 0; i < n; i++) {
            if (middle == nums[i]) {
                continue;
            } else if (partener.contains(nums[i])) {
                partener.remove(nums[i]);
            } else {
                partener.add(n - nums[i]);
            }
        }
        if (1 == partener.size()) {
            return partener.iterator().next();
        } else {
            return middle;
        }
    }

    public static void main(String[] args) {
        MissingNumber mn = new MissingNumber();

        int[][] tests = {
            {0, 2},
            {3, 0, 1},
            {9, 6, 4, 2, 3, 5, 7, 0, 1},
        };

        for (int[] test: tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(mn.missingNumber(test));
            System.out.println();
        }
    }
}

