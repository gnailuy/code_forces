// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/

import java.util.List;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) {
        val = x;
        left = right = null;
    }
}

class SortedArrayToBST {
    public TreeNode sortedArrayToBST(int[] nums, int left, int right) {
        if (left > right) {
            return null;
        }

        int middle = (left + right) / 2;
        TreeNode root = new TreeNode(nums[middle]);
        root.left = sortedArrayToBST(nums, left, middle-1);
        root.right = sortedArrayToBST(nums, middle+1, right);

        return root;
    }

    public TreeNode sortedArrayToBST(int[] nums) {
        return sortedArrayToBST(nums, 0, nums.length-1);
    }

    public static void main(String[] args) {
        SortedArrayToBST sabst = new SortedArrayToBST();
        BTLevelOrderTraversal btlot = new BTLevelOrderTraversal();

        int[][] tests = {
            {},
            {1},
            {1, 1},
            {1, 1, 1},
            {1, 2, 3},
            {1, 2, 2, 3, 4, 4, 5},
            {-10, -3, 0, 5, 9},
        };
        for (int[] test : tests) {
            TreeNode root = sabst.sortedArrayToBST(test);

            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            List<List<Integer>> lo = btlot.levelOrder(root);
            for (List<Integer> l: lo) {
                for (Integer i: l) {
                    System.out.print(i + " ");
                }
                System.out.println();
            }
            System.out.println();
        }
    }
}
