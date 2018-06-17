// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/565/

class NumberOf1Bits {
    public int hammingWeight(int n) {
        int n1 = 0;
        int mask = 1;
        for (int i = 0; i < 32; i++) {
            if ((n&mask) != 0) n1++;
            mask = mask << 1;
        }
        return n1;
    }

    public static void main(String[] args) {
        NumberOf1Bits no1 = new NumberOf1Bits();

        int[] tests = {
            11,
            128,
        };

        for (int test: tests) {
            System.out.println(test);
            System.out.println(no1.hammingWeight(test));
            System.out.println();
        }
    }
}

