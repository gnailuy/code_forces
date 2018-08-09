// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/784/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class OddEvenLinkedList{
    public ListNode oddEvenList(ListNode head) {
        if (null == head || null == head.next) {
            return head;
        }
        ListNode odd = head; // Last Old Node
        ListNode even = head.next; // Last Even Node

        while (null != odd && null != even) {
            if (null == even.next) {
                break;
            } else {
                ListNode tmp = odd.next; // Fist Even Node
                odd.next = even.next;
                even.next = odd.next.next;
                odd.next.next = tmp;
            }

            odd = odd.next;
            even = even.next;
        }

        return head;
    }

    public static void main(String[] args) {
        OddEvenLinkedList oel = new OddEvenLinkedList();

        int[][] tests = {
            {1, 2, 3, 4},
            {1, 2, 3, 4, 5},
            {2, 1, 3, 5, 6, 4, 7},
            {1},
            {1, 2},
            {},
        };
        for (int[] test : tests) {
            ListNode head = null;
            ListNode p = null;
            for (int i = 0; i < test.length; i++) {
                ListNode t = new ListNode(test[i]);
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

            head = oel.oddEvenList(head);

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
        }
    }
}
