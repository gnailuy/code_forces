// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/789/

public class TreeLinkNextRightPointers {
    public void connect(TreeLinkNode root) {
        TreeLinkNode layer = root;
        TreeLinkNode nextLayer = null;
        while (null != layer && null != layer.left) {
            nextLayer = layer.left;
            while (null != layer) {
                layer.left.next = layer.right;
                if (null != layer.next) layer.right.next = layer.next.left;
                layer = layer.next;
            }
            layer = nextLayer;
        }
    }

    public static void main(String[] args) {
        TreeLinkNextRightPointers tlnrp = new TreeLinkNextRightPointers();

        int[][] tests = {
            {0, 1, 2, 3, 4, 5, 6},
        };
        for (int[] test : tests) {
            // TODO: Build Tree
        }
    }
}
