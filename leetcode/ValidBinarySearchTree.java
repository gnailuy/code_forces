// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/

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

class ValidBinarySearchTree {
    public int max(TreeNode root) {
        if (null == root.right) {
            return root.val;
        } else {
            return max(root.right);
        }
    }

    public int min(TreeNode root) {
        if (null == root.left) {
            return root.val;
        } else {
            return min(root.left);
        }
    }

    public boolean isValidBST(TreeNode root) {
        if (null == root) return true;

        if (!isValidBST(root.left)) return false;
        if (!isValidBST(root.right)) return false;
        if (null != root.left) {
            if (max(root.left) >= root.val) return false;
        }
        if (null != root.right) {
            if (min(root.right) <= root.val) return false;
        }

        return true;
    }

    public static void main(String[] args) {
        ValidBinarySearchTree vbst = new ValidBinarySearchTree();

        int[][] tests = {
            {},
            {1},
            {1, 2},
            {1, 2, 3},
            {2, 1, 3},
            {4, 2, 6, 1, 3, 5, 7},
            {5, 1, 4, -1, -1, 3, 6},
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

            System.out.println(vbst.isValidBST(root));
        }
    }
}
