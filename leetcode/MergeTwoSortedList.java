// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/771/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class MergeTwoSortedList{
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode head = null;
        ListNode p = null;

        while (null != l1 && null != l2) {
            while (null != l1 && null != l2 && l1.val <= l2.val) {
                if (null == head) {
                    head = new ListNode(l1.val);
                    p = head;
                } else {
                    p.next = new ListNode(l1.val);
                    p = p.next;
                }
                l1 = l1.next;
            }
            while (null != l1 && null != l2 && l2.val <= l1.val) {
                if (null == head) {
                    head = new ListNode(l2.val);
                    p = head;
                } else {
                    p.next = new ListNode(l2.val);
                    p = p.next;
                }
                l2 = l2.next;
            }
        }
        while (null != l1) {
            if (null == head) {
                head = new ListNode(l1.val);
                p = head;
            } else {
                p.next = new ListNode(l1.val);
                p = p.next;
            }
            l1 = l1.next;
        }
        while (null != l2) {
            if (null == head) {
                head = new ListNode(l2.val);
                p = head;
            } else {
                p.next = new ListNode(l2.val);
                p = p.next;
            }
            l2 = l2.next;
        }

        return head;
    }

    public static void main(String[] args) {
        MergeTwoSortedList msl = new MergeTwoSortedList();

        int[][][] tests = {
            {{1, 2, 4}, {1, 3, 4}},
            {{1, 3, 5}, {2, 4, 6}},
            {{1, 2, 3}, {}},
            {{1}, {}},
            {{}, {1}},
            {{}, {}},
        };
        for (int[][] test : tests) {
            ListNode[] l = new ListNode[2];
            for (int j = 0; j < 2; j++) {
                ListNode p = null;
                for (int i = 0; i < test[j].length; i++) {
                    ListNode t = new ListNode(test[j][i]);
                    if (null == l[j]) {
                        l[j] = t;
                    }
                    if (null != p) {
                        p.next = t;
                    }
                    p = t;
                }

                p = l[j];
                while (null != p) {
                    System.out.print(p.val);
                    if (null != p.next) {
                        System.out.print("->");
                    }
                    p = p.next;
                }
                System.out.print(" ");
            }
            System.out.println();

            ListNode head = msl.mergeTwoLists(l[0], l[1]);

            ListNode p = head;
            while (null != p) {
                System.out.print(p.val + " ");
                p = p.next;
            }
            System.out.println();
        }
    }
}
