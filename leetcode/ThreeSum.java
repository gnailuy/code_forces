// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/

import java.util.Arrays;
import java.util.List;
import java.util.LinkedList;

class ThreeSum {
    private int binary_search(int target, int[] nums, int low, int high) {
        if (low > high) return -1;

        int middle = (low+high)/2;
        if (target == nums[middle]) {
            return middle;
        } else if (target < nums[middle]) {
            return binary_search(target, nums, low, middle-1);
        } else {
            return binary_search(target, nums, middle+1, high);
        }
    }

    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        List<List<Integer>> result = new LinkedList<List<Integer>>();
        for (int i = 0; i < n-2; i++) {
            if (i > 0 && nums[i] == nums[i-1]) continue;
            for (int j = i+1; j < n-1; j++) {
                if (j > i+1 && nums[j] == nums[j-1]) continue;
                int target = 0 - nums[i] - nums[j];
                if (target < nums[j+1] || target > nums[n-1]) continue;
                int k = binary_search(target, nums, j+1, n-1);
                if (k > 0) {
                    List<Integer> l = new LinkedList<Integer>();
                    l.add(nums[i]);
                    l.add(nums[j]);
                    l.add(nums[k]);
                    result.add(l);
                }
            }
        }
        return result;
    }

    public static void main(String[] args) {
        ThreeSum ts = new ThreeSum();

        int[][] tests = {
            {-1, 0, 1, 2, -1, -4},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();

            List<List<Integer>> result = ts.threeSum(test);
            System.out.println(result.size());
            for (List<Integer> l: result) {
                for (Integer i: l) {
                    System.out.print(i + " ");
                }
                System.out.println();
            }
            System.out.println();
        }
    }
}
