// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/773/

import java.util.HashSet;

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) {
        val = x;
        next = null;
    }
}

class LinkedListCycle{
    public boolean hasCycle(ListNode head) {
        ListNode p = head;
        HashSet<ListNode> met = new HashSet<ListNode>();
        while (null != p) {
            if (met.contains(p)) {
                return true;
            } else {
                met.add(p);
            }
            p = p.next;
        }
        return false;
    }

    public static void main(String[] args) {
        LinkedListCycle llc = new LinkedListCycle();

        int[][][] tests = {
            {{1, 2, 3, 4, 5}, {2}},
        };
        for (int[][] test : tests) {
            ListNode head = null;
            ListNode p = null;
            ListNode q = null;
            for (int i = 0; i < test[0].length; i++) {
                ListNode t = new ListNode(test[0][i]);
                if (null == head) {
                    head = t;
                }
                if (null != p) {
                    p.next = t;
                }
                p = t;
                if (test[1][0] == i) {
                    q = t;
                }
            }
            System.out.println(llc.hasCycle(head));

            p.next = q; // Point the last one to q
            p = head;
            for (int i = 0; i <= test[0].length; i++) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
            System.out.println(llc.hasCycle(head));
        }
    }
}
