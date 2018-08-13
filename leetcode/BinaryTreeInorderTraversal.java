// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/786/

import java.util.LinkedList;
import java.util.List;

public class BinaryTreeInorderTraversal {
    public void inorderHelper(TreeNode root, List<Integer> out) {
        if (null == root) return;

        if (null != root.left) {
            inorderHelper(root.left, out);
        }
        out.add(root.val);
        if (null != root.right) {
            inorderHelper(root.right, out);
        }
    }

    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new LinkedList<Integer>();
        inorderHelper(root, result);
        return result;
    }

    public static void main(String[] args) {
        BinaryTreeInorderTraversal btit = new BinaryTreeInorderTraversal();

        int[][] tests = {
            {1, -1, 2, 3},
        };
        for (int[] test : tests) {
            // TODO: Build Tree
            // TODO: Inorder Traversal
        }
    }
}
