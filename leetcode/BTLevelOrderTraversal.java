// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/628/

import java.util.List;
import java.util.LinkedList;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) {
        val = x;
        left = right = null;
    }
}

class BTLevelOrderTraversal {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> lo = new LinkedList<List<Integer>>();

        if (null != root) {
            List<TreeNode> level = new LinkedList<TreeNode>();
            List<TreeNode> next = null;
            List<Integer> l = null;

            level.add(root);

            while (!level.isEmpty()) {
                next = new LinkedList<TreeNode>();
                l = new LinkedList<Integer>();
                for (int i = 0; i < level.size(); i++) {
                    TreeNode node = level.get(i);
                    l.add(node.val);
                    if (null != node.left) next.add(node.left);
                    if (null != node.right) next.add(node.right);
                }
                lo.add(l);
                level = next;
            }
        }

        return lo;
    }

    public static void main(String[] args) {
        BTLevelOrderTraversal btlot = new BTLevelOrderTraversal();

        int[][] tests = {
            {},
            {1},
            {1, 1},
            {1, 1, 1},
            {1, 2, 1},
            {1, 2, 2, 3, 4, 4, 3},
            {1, 2, 2, -1, 3, -1, 3},
            {3, 9, 20, -1, -1, 15, 7},
        };
        for (int[] test : tests) {
            TreeNode root = null;
            LinkedList<TreeNode> level = null;
            LinkedList<TreeNode> next = null;

            if (test.length > 0) {
                root = new TreeNode(test[0]);
                level = new LinkedList<TreeNode>();
                level.add(root);
                for (int i = 2; ; i *= 2) {
                    boolean end = false;
                    next = new LinkedList<TreeNode>();
                    for (int j = i; j < 2*i; j++) {
                        int k = j-1;
                        if (k < test.length) {
                            TreeNode p = null;
                            if (-1 != test[k]) {
                                p = new TreeNode(test[k]);
                            }
                            if (j%2 == 0) {
                                level.get((j-i)/2).left = p;
                            } else {
                                level.get((j-i)/2).right = p;
                            }
                            next.add(p);
                        } else {
                            end = true;
                            break;
                        }
                    }
                    level = next;
                    if (end) break;
                }

                level = new LinkedList<TreeNode>();
                level.add(root);
                while (!level.isEmpty()) {
                    TreeNode p = level.poll();
                    System.out.print(p.val + " ");
                    if (null != p.left) {
                        level.add(p.left);
                    }
                    if (null != p.right) {
                        level.add(p.right);
                    }
                }
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
