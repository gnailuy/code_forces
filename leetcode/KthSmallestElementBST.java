// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/790/

public class KthSmallestElementBST {
    class RankValue {
        int rank;
        int value;

        public RankValue() {
            rank = 0;
            value = Integer.MIN_VALUE;
        }

        public RankValue(int rank) {
            this.rank = rank;
            value = Integer.MIN_VALUE;
        }

        public void setRank(int n) {
            rank = n;
        }

        public void setValue(int v) {
            value = v;
        }
    }

    public RankValue inorderHelper(TreeNode root, int k, RankValue rv) {
        if (null == root) return rv;

        RankValue result = inorderHelper(root.left, k, rv);
        if (result.rank == k) return result;

        result.rank += 1;
        if (k == result.rank) {
            result.value = root.val;
            return result;
        }

        return inorderHelper(root.right, k, result);
    }

    public int kthSmallest(TreeNode root, int k) {
        RankValue result = inorderHelper(root, k, new RankValue());
        return result.value;
    }

    public static void main(String[] args) {
        KthSmallestElementBST kbst = new KthSmallestElementBST();

        int[][][] tests = {
            {{3, 1, 4, -1, 2}, {1}},
            {{5, 3, 6, 2, 4, -1, -1, 1}, {3}},
        };
        for (int[][] test : tests) {
            // TODO: Build Tree
        }
    }
}
