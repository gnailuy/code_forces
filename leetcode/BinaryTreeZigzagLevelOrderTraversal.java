// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/787/

import java.util.Deque;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class BinaryTreeZigzagLevelOrderTraversal {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new LinkedList<List<Integer>>();
        if (null == root) return result;

        Deque<TreeNode> level = new LinkedList<TreeNode>();
        level.add(root);

        List<Integer> values = new LinkedList<Integer>();
        for (TreeNode node: level) {
            values.add(node.val);
        }
        result.add(values);

        int n = 2;
        while (level.size() > 0) {
            Deque<TreeNode> newLevel = new LinkedList<TreeNode>();

            while (level.size() > 0) {
                TreeNode next = level.poll();
                if (null != next) {
                    if (next.left != null) newLevel.add(next.left);
                    if (next.right != null) newLevel.add(next.right);
                }
            }

            if (newLevel.size() <= 0) {
                break;
            }
            values = new LinkedList<Integer>();
            if (n%2 == 0) {
                Iterator<TreeNode> it = newLevel.descendingIterator();
                while (it.hasNext()) {
                    TreeNode next = it.next();
                    values.add(next.val);
                }
            } else {
                for (TreeNode node: newLevel) {
                    values.add(node.val);
                }
            }
            result.add(values);

            level = newLevel;
            n++;
        }

        return result;
    }

    public static void main(String[] args) {
        BinaryTreeZigzagLevelOrderTraversal btzt = new BinaryTreeZigzagLevelOrderTraversal();

        int[][] tests = {
            {3, 9, 20, -1, -1, 15, 7}
        };
        for (int[] test : tests) {
            // TODO: Build Tree
            // TODO: ZigzagLevelOrder Traversal
        }
    }
}
