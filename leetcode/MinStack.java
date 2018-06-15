// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/562/

class MinStack {
    class Node {
        int num;
        Node next;

        public Node(int n) {
            num = n;
            next = null;
        }
    }

    Node head;
    Node min;

    /** initialize your data structure here. */
    public MinStack() {
        head = null;
        min = null;
    }

    public void push(int x) {
        Node xn = new Node(x);
        xn.next = head;
        head = xn;

        Node nn;
        if (null == min || min.num > x) {
            nn = new Node(x);
        } else {
            nn = new Node(min.num);
        }
        nn.next = min;
        min = nn;
    }

    public void pop() {
        if (null != head) head = head.next;
        if (null != min) min = min.next;
    }

    public int top() {
        assert null != head;
        return head.num;
    }

    public int getMin() {
        assert null != min;
        return min.num;
    }

    public static void main(String[] args) {
        int[][] tests = {
            {1, 2, 3},
            {1, 2, 3, 4, 5},
            {-2, 1, -3, 4, -1, 2, 1, -5, 4},
            {-2, -1, -3, -4, -6, -7, -5},
            {-2, -1, -3, -4, 1, -6, -7, -5},
        };
        for (int[] test : tests) {
            MinStack ms = new MinStack();
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            for (int i = 0; i < test.length; i++) {
                ms.push(test[i]);
            }
            System.out.println(ms.getMin());
            ms.pop();
            ms.pop();
            System.out.println(ms.top());
            System.out.println();
        }
    }
}
