// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/783/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class AddTwoNumbers {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode result = null;
        ListNode p = null;
        int carry = 0;

        while (null != l1 || null != l2) {
            int digit = carry;
            if (null != l1) {
                digit += l1.val;
            }
            if (null != l2) {
                digit += l2.val;
            }
            if (digit >= 10) {
                carry = 1;
                digit -= 10;
            } else {
                carry = 0;
            }

            ListNode newNode = new ListNode(digit);
            if (null == result) {
                result = newNode;
            }
            if (null == p) {
                p = newNode;
            } else {
                p.next = newNode;
                p = newNode;
            }

            if (null != l1) l1 = l1.next;
            if (null != l2) l2 = l2.next;
        }

        if (carry > 0) {
            ListNode newNode = new ListNode(carry);
            p.next = newNode;
            p = newNode;
        }

        return result;
    }

    public static void main(String[] args) {
        AddTwoNumbers atn = new AddTwoNumbers();

        int[][][] tests = {
            {{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}},
            {{2, 4, 3}, {5, 6, 4}},
            {{9, 9, 9}, {9, 9, 9, 9, 9}},
            {{9, 9, 9, 9, 9}, {9, 9, 9}},
        };
        for (int[][] test : tests) {
            ListNode[] head = new ListNode[2];
            ListNode p = null;
            for (int j = 0; j < 2; j++) {
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
            }

            for (int j = 0; j < 2; j++) {
                p = head[j];
                while (null != p) {
                    System.out.print(p.val + " ");
                    p = p.next;
                }
                System.out.println();
            }

            ListNode result = atn.addTwoNumbers(head[0], head[1]);

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
