// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/560/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class ReverseLinkedList{
    public ListNode reverseList(ListNode head) {
        if (null == head) return head;

        ListNode p = head.next;
        head.next = null;
        while (null != p) {
            ListNode q = p.next;
            p.next = head;
            head = p;
            p = q;
        }
        return head;
    }

    public static void main(String[] args) {
        ReverseLinkedList rll = new ReverseLinkedList();

        int[][] tests = {
            {1, 2, 3, 4, 5},
            {1},
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

            head = rll.reverseList(head);

            p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
        }
    }
}
