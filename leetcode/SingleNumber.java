// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/

class SingleNumber {
    public int singleNumber(int[] nums) {
        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            result ^= nums[i];
        }
        return result;
    }

    public static void main(String[] args) {
        SingleNumber sn = new SingleNumber();

        int[][] tests = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(sn.singleNumber(test));
        }
    }
}
