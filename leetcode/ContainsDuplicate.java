// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/

import java.util.Random;

class ContainsDuplicate {
    public void shuffle(int[] nums) {
        Random rand = new Random();
        for (int i = 0; i < nums.length-1; i++) {
            int n = rand.nextInt(nums.length - i - 1) + 1 + i;
            int tmp = nums[i];
            nums[i] = nums[n];
            nums[n] = tmp;
        }
    }

    public boolean containsDuplicate(int[] nums, int lo, int hi) {
        if (hi <= lo) {
            return false;
        }

        int pivot = nums[lo];
        int i = lo+1, j = hi;
        while (i <= j) {
            while (i <= hi && nums[i] < pivot) i++;
            while (j >= lo+1 && nums[j] > pivot) j--;
            if (i > j) {
                break;
            } else if (nums[i] == pivot || nums[j] == pivot) {
                return true;
            } else {
                int tmp = nums[i];
                nums[i] = nums[j];
                nums[j] = tmp;
            }
        }

        if (i > hi || j <= lo) {
            return containsDuplicate(nums, lo+1, hi);
        } else {
            boolean left = containsDuplicate(nums, lo+1, j);
            boolean right = containsDuplicate(nums, i, hi);
            return left || right;
        }
    }

    public boolean containsDuplicate(int[] nums) {
        if (nums.length <= 1) {
            return false;
        }

        shuffle(nums);
        return containsDuplicate(nums, 0, nums.length-1);
    }

    public static void main(String[] args) {
        ContainsDuplicate cd = new ContainsDuplicate();

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
