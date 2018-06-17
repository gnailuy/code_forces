// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/648/

class ReverseBits {
    public int reverseBits(int n) {
        int result = 0;
        int mask = 1;
        for (int i = 0; i < 32; i++) {
            if ((n&mask) != 0) {
                result += (1<<(31-i));
            }
            mask <<= 1;
        }
        return result;
    }

    public static void main(String[] args) {
        ReverseBits rb = new ReverseBits();

        int[] tests = {
            11,
            128,
            43261596,
        };

        for (int test: tests) {
            System.out.println(test);
            System.out.println(rb.reverseBits(test));
            System.out.println();
        }
    }
}

