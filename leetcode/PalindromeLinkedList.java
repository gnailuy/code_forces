// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/772/

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class PalindromeLinkedList{
    public boolean isPalindrome(ListNode head) {
        ListNode p = head;
        int length = 0;
        while (null != p) {
            length++;
            p = p.next;
        }
        if (length <= 1) {
            return true;
        }

        int[] firstHalf = new int[length/2];
        p = head;
        for (int i = 0; i < length/2; i++) {
            firstHalf[i] = p.val;
            p = p.next;
        }
        if (length % 2 != 0) {
            p = p.next;
        }
        for (int i = length/2-1; i >= 0; i--) {
            if (firstHalf[i] != p.val) {
                return false;
            }
            p = p.next;
        }
        return true;
    }

    public static void main(String[] args) {
        PalindromeLinkedList pll = new PalindromeLinkedList();

        int[][] tests = {
            {1, 2},
            {1, 2, 2, 1},
            {1, 1},
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
            System.out.println(pll.isPalindrome(head));
        }
    }
}
