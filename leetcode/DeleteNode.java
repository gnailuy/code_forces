// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/553/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class DeleteNode {
    private static ListNode head = null;

    public void deleteNode(ListNode node) {
        node.val = node.next.val;
        node.next = node.next.next;
    }

    public static void main(String[] args) {
        DeleteNode dn = new DeleteNode();

        int[][][] tests = {
            {{1, 2, 3, 4}, {3}},
        };
        for (int[][] test : tests) {
            ListNode p = null, target = null;
            for (int i = 0; i < test[0].length; i++) {
                ListNode t = new ListNode(test[0][i]);
                if (null == head) {
                    head = t;
                }
                if (null != p) {
                    p.next = t;
                }
                p = t;
                if (null == target && test[0][i] == test[1][0]) {
                    target = t;
                }
            }

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
            System.out.println(target.val);

            dn.deleteNode(target);

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
        }
    }
}
