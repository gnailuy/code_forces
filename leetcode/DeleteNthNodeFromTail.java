// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/603/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class DeleteNthNodeFromTail {
    private static ListNode head = null;

    public ListNode removeNthFromEnd(ListNode head, int n) {
        assert n > 0;

        int c = n;
        ListNode p = null; // Before to_delete
        ListNode t = head;
        while (null != t.next) {
            t = t.next;
            c--;
            if (0 == c) {
                p = head;
            } else if (c < 0) {
                p = p.next;
            }
        }
        assert c <= 1;

        if (null == p) {
            return head.next;
        } else {
            p.next = p.next.next;
            return head;
        }
    }

    public static void main(String[] args) {
        DeleteNthNodeFromTail dne = new DeleteNthNodeFromTail();

        int[][][] tests = {
            {{1, 2, 3, 4, 5}, {1}},
            {{1, 2, 3, 4, 5}, {2}},
            {{1, 2, 3, 4, 5}, {3}},
            {{1, 2, 3, 4, 5}, {4}},
            {{1, 2, 3, 4, 5}, {5}},
            {{1, 2}, {1}},
            {{1, 2}, {2}},
            {{1}, {1}},
        };
        for (int[][] test : tests) {
            ListNode p = null;
            int n = test[1][0];
            for (int i = 0; i < test[0].length; i++) {
                ListNode t = new ListNode(test[0][i]);
                if (null == head) {
                    head = t;
                }
                if (null != p) {
                    p.next = t;
                }
                p = t;
            }

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
            System.out.println(n);

            head = dne.removeNthFromEnd(head, n);

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println("\n");

            head = null;
        }
    }
}
