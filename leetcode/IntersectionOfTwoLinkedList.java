// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/785/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class IntersectionOfTwoLinkedList {
    public int findLength(ListNode head) {
        int len = 0;
        ListNode p = head;
        while (null != p) {
            p = p.next;
            len++;
        }
        return len;
    }

    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        int la = findLength(headA);
        int lb = findLength(headB);

        ListNode p = headA;
        ListNode q = headB;
        while (la < lb) {
            q = q.next;
            lb--;
        }
        while (la > lb) {
            p = p.next;
            la--;
        }

        while (null != p && null != q) {
            if (p == q) {
                return p;
            }

            p = p.next;
            q = q.next;
        }

        return null;
    }

    public static void main(String[] args) {
        IntersectionOfTwoLinkedList iot = new IntersectionOfTwoLinkedList();

        int[][][] tests = {
            {{1, 2}, {3, 4, 5}, {6, 7, 8}},
            {{1, 2, 3}, {4, 5, 6}, {}},
            {{1}, {}, {2, 3, 4}},
            {{}, {}, {1, 2, 3}},
            {{}, {}, {}},
        };
        for (int[][] test : tests) {
            ListNode[] head = new ListNode[3];
            ListNode[] tail = new ListNode[3];
            ListNode p = null;
            for (int j = 0; j < 3; j++) {
                p = null;
                for (int i = 0; i < test[j].length; i++) {
                    ListNode t = new ListNode(test[j][i]);
                    if (null == head[j]) {
                        head[j] = t;
                    }
                    if (null != p) {
                        p.next = t;
                    }
                    p = t;
                }
                tail[j] = p;
            }

            for (int j = 0; j < 2; j++) {
                if (null != tail[j]) {
                    tail[j].next = head[2];
                } else {
                    head[j] = head[2];
                }
            }

            for (int j = 0; j < 2; j++) {
                p = head[j];
                while (null != p) {
                    System.out.print(p.val + " ");
                    p = p.next;
                }
                System.out.println();
            }

            ListNode result = iot.getIntersectionNode(head[0], head[1]);

            p = result;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
            System.out.println();
        }
    }
}
