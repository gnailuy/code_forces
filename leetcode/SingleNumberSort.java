// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/

import java.util.Arrays;

class SingleNumberSort {
    public int singleNumber(int[] nums) {
        Arrays.sort(nums);
        for (int i = 0; i < nums.length-1; i += 2) {
            if (nums[i] != nums[i+1]) {
                return nums[i];
            }
        }
        return nums[nums.length-1];
    }

    public static void main(String[] args) {
        SingleNumberSort sn = new SingleNumberSort();

        int[][] tests = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(sn.singleNumber(test));
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
        }
    }
}
