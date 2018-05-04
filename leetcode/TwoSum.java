// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/

import java.util.HashMap;

class TwoSum {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> store = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            int other = target - nums[i];
            if (store.containsKey(other)) {
                return new int[]{store.get(other), i};
            } else {
                store.put(nums[i], i);
            }
        }
        return new int[]{-1, -1}; // Not found
    }

    public static void main(String[] args) {
        TwoSum ts = new TwoSum();

        int[][][] tests = {
            {{2, 7, 11, 15}, {9}},
        };
        for (int[][] test : tests) {
            System.out.println(test[0].length + " " + test[1][0]);
            for (int i = 0; i < test[0].length; i++) {
                System.out.print(test[0][i] + " ");
            }
            System.out.println();
            int[] result = ts.twoSum(test[0], test[1][0]);
            System.out.println(result.length);
            for (int i = 0; i < result.length; i++) {
                System.out.print(result[i] + " ");
            }
            System.out.println();
        }
    }
}
