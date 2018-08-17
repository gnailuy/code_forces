// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/795/

import java.util.TreeSet;
import java.util.LinkedList;
import java.util.List;

class Permutations {
    private void permuteHelper(TreeSet<Integer> nums, LinkedList<Integer> current,
            List<List<Integer>> result) {
        if (nums.isEmpty()) {
            List<Integer> permute = new LinkedList<Integer>();
            for (Integer i: current) permute.add(i);
            result.add(permute);
        }

        TreeSet<Integer> numsCopy = new TreeSet<Integer>();
        for (Integer i: nums) numsCopy.add(i);

        for (Integer i: nums) {
            current.add(i);
            numsCopy.remove(i);

            permuteHelper(numsCopy, current, result);

            current.removeLast();
            numsCopy.add(i);
        }
    }

    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new LinkedList<List<Integer>>();

        TreeSet<Integer> numSet = new TreeSet<Integer>();
        for (int i = 0; i < nums.length; i++) {
            numSet.add(nums[i]);
        }

        permuteHelper(numSet, new LinkedList<Integer>(), result);
        return result;
    }

    public static void main(String[] args) {
        Permutations p = new Permutations();

        int[][] tests = {
            {1, 2, 3},
            {1, 2, 3, 4},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();

            List<List<Integer>> result = p.permute(test);
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
