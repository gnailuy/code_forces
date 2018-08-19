// https://leetcode.com/explore/interview/card/top-interview-questions-medium/112/design/812/

import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class SerializeDeserializeBinaryTree {
    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (null == root) {
            return "";
        }

        StringBuilder sb = new StringBuilder();
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.add(root);
        sb.append(root.val + ",");

        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            if (null == node.left) {
                sb.append(",");
            } else {
                sb.append(node.left.val + ",");
                queue.add(node.left);
            }
            if (null == node.right) {
                sb.append(",");
            } else {
                sb.append(node.right.val + ",");
                queue.add(node.right);
            }
        }

        int i = sb.length()-1;
        while (i >= 0 && sb.charAt(i) == ',') sb.deleteCharAt(i--);
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if ("".equals(data)) return null;

        String[] values = data.split(",");
        TreeNode result = new TreeNode(Integer.parseInt(values[0]));
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.add(result);

        int i = 1;
        while (!queue.isEmpty() && i < values.length) {
            TreeNode node = queue.poll();
            if (!"".equals(values[i])) {
                int num = Integer.parseInt(values[i]);
                node.left = new TreeNode(num);
                queue.add(node.left);
            }
            if (i+1 < values.length && !"".equals(values[i+1])) {
                int num = Integer.parseInt(values[i+1]);
                node.right = new TreeNode(num);
                queue.add(node.right);
            }
            i += 2;
        }

        return result;
    }

    public static void main(String[] args) {
        SerializeDeserializeBinaryTree codec = new SerializeDeserializeBinaryTree();
        BTLevelOrderTraversal btlot = new BTLevelOrderTraversal();

        String[] tests = {
            "1,2,3,,,4,5",
            "1,2,3,,,4,5,,6",
            "1,2,3,,,4,5,,6,7,8",
            "",
        };

        for (String test: tests) {
            TreeNode root = codec.deserialize(test);
            List<List<Integer>> lo = btlot.levelOrder(root);
            for (List<Integer> l: lo) {
                for (Integer i: l) {
                    System.out.print(i + " ");
                }
                System.out.println();
            }
            System.out.println();
            System.out.println(codec.serialize(root));
        }
    }
}

