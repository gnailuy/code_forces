// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/

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

class MaximumDepthBinaryTree {
    public int maxDepth(TreeNode root) {
        if (null == root) return 0;

        int height = 1;
        LinkedList<TreeNode> next = null;
        LinkedList<TreeNode> current = new LinkedList<TreeNode>();
        current.add(root);
        while (true) {
            next = new LinkedList<TreeNode>();
            while (!current.isEmpty()) {
                TreeNode p = current.poll();
                if (null != p.left) {
                    next.add(p.left);
                }
                if (null != p.right) {
                    next.add(p.right);
                }
            }
            if (next.size() == 0) {
                break;
            }
            height++;
            current = next;
        }

        return height;
    }

    public static void main(String[] args) {
        MaximumDepthBinaryTree mdbt = new MaximumDepthBinaryTree();

        int[][] tests = {
            {},
            {1},
            {1, 1},
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

            System.out.println(mdbt.maxDepth(root));
        }
    }
}
