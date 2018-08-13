// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/

import java.util.LinkedList;

class SymmetricTree {
    public boolean equalNode(TreeNode a, TreeNode b) {
        if (null == a && null == b) {
            return true;
        } else if (null == a || null == b) {
            return false;
        } else {
            return a.val == b.val;
        }
    }

    public boolean isSymmetric(TreeNode root) {
        if (null == root) {
            return true;
        } else if (null == root.left && null == root.right) {
            return true;
        } else if (null == root.left || null == root.right) {
            return false;
        }

        LinkedList<TreeNode> level = new LinkedList<TreeNode>();
        LinkedList<TreeNode> next = null;
        level.add(root.left);
        level.add(root.right);

        while (!level.isEmpty()) {
            next = new LinkedList<TreeNode>();
            for (int i = 0; i < level.size()/2; i++) {
                if (!equalNode(level.get(i), level.get(level.size()-i-1))) {
                    return false;
                }
                if (null != level.get(i)) {
                    next.add(level.get(i).left);
                    next.add(level.get(i).right);
                }
            }
            for (int i = level.size()/2; i < level.size(); i++) {
                if (null != level.get(i)) {
                    next.add(level.get(i).left);
                    next.add(level.get(i).right);
                }
            }
            level = next;
        }

        return true;
    }

    public static void main(String[] args) {
        SymmetricTree st = new SymmetricTree();

        int[][] tests = {
            {},
            {1},
            {1, 1},
            {1, 1, 1},
            {1, 2, 1},
            {1, 2, 2, 3, 4, 4, 3},
            {1, 2, 2, -1, 3, -1, 3},
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

            System.out.println(st.isSymmetric(root));
        }
    }
}
