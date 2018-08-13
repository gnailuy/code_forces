// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/788/

import java.util.List;

public class ConstructBinaryTree {
    public int indexOf(int[] array, int target) {
        for (int i = 0; i < array.length; i++) {
            if (array[i] == target) return i;
        }
        return -1;
    }

    public TreeNode buildTree(int[] preorder, int[] inorder, int rootIndex, int inLeft, int inRight) {
        if (rootIndex >= preorder.length) return null;
        if (inLeft >= inRight) return null;

        TreeNode root = new TreeNode(preorder[rootIndex]);
        int inorderRootIndex = indexOf(inorder, root.val);

        root.left = buildTree(preorder, inorder, rootIndex+1, inLeft, inorderRootIndex);
        root.right = buildTree(preorder, inorder, rootIndex+inorderRootIndex-inLeft+1, inorderRootIndex+1, inRight);

        return root;
    }

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return buildTree(preorder, inorder, 0, 0, inorder.length);
    }

    public static void main(String[] args) {
        ConstructBinaryTree cbt = new ConstructBinaryTree();
        BinaryTreeInorderTraversal btit = new BinaryTreeInorderTraversal();

        int[][][] tests = {
            {{3, 9, 20, 15, 7}, {9, 3, 15, 20, 7}},
        };
        for (int[][] test : tests) {
            TreeNode tree = cbt.buildTree(test[0], test[1]);
            List<Integer> result = btit.inorderTraversal(tree);
            for (Integer i: result) {
                System.out.print(i + " ");
            }
            System.out.println();
        }
    }
}
