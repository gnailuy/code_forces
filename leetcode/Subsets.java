// https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/796/

import java.util.LinkedList;
import java.util.List;

class Subsets {
    private void subsetsHelper(int[] nums, int index,
            LinkedList<Integer> current, List<List<Integer>> result) {
        List<Integer> copy = new LinkedList<Integer>();
        for (Integer i: current) { copy.add(i); }
        result.add(copy);

        for (int i = index; i < nums.length; i++) {
            current.add(nums[i]);
            subsetsHelper(nums, i+1, current, result);
            current.removeLast();
        }
    }

    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new LinkedList<List<Integer>>();

        LinkedList<Integer> current = new LinkedList<Integer>();
        subsetsHelper(nums, 0, current, result);

        return result;
    }

    public static void main(String[] args) {
        Subsets ss = new Subsets();

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

            List<List<Integer>> result = ss.subsets(test);
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
