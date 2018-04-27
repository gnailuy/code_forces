// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/

import java.util.HashSet;

class ContainsDuplicateHashSet {
    public boolean containsDuplicate(int[] nums) {
        if (nums.length <= 1) {
            return false;
        }

        HashSet<Integer> set = new HashSet<Integer>();
        for (int i = 0; i < nums.length; i++) {
            if (set.contains(nums[i])) {
                return true;
            } else {
                set.add(nums[i]);
            }
        }

        return false;
    }

    public static void main(String[] args) {
        ContainsDuplicateHashSet cd = new ContainsDuplicateHashSet();

        int[][] tests = {
            {1},
            {1, 1, 2},
            {0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
            {0, 1, 2, 3, 4, 5},
            {4, 1, 0, 3, 6, 2, 5}
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(cd.containsDuplicate(test));
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
