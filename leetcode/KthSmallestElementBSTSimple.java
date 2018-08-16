// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/790/

public class KthSmallestElementBSTSimple {
    class Rank {
        int rank;
        public Rank(int n) {this.rank = n;}
    }

    public int inorderHelper(TreeNode root, int k, Rank rank) {
        if (null == root) return Integer.MIN_VALUE;

        if (null != root.left) {
            int result = inorderHelper(root.left, k, rank);
            if (result > Integer.MIN_VALUE) return result;
        }

        rank.rank += 1;
        if (k == rank.rank) return root.val;

        return inorderHelper(root.right, k, rank);
    }

    public int kthSmallest(TreeNode root, int k) {
        return inorderHelper(root, k, new Rank(0));
    }

    public static void main(String[] args) {
        KthSmallestElementBSTSimple kbst = new KthSmallestElementBSTSimple();

        int[][][] tests = {
            {{3, 1, 4, -1, 2}, {1}},
            {{5, 3, 6, 2, 4, -1, -1, 1}, {3}},
        };
        for (int[][] test : tests) {
            // TODO: Build Tree
        }
    }
}
