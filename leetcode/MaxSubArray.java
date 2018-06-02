// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/566/

class MaxSubArray {
    public int maxSubArray(int[] nums) {
        int maxMsa = 0;
        int currentMsa = 0;
        int maxNeg = Integer.MIN_VALUE;
        boolean allNeg = true;

        int i = 0;
        while (i < nums.length && nums[i] < 0) {
            if (maxNeg < nums[i]) maxNeg = nums[i];
            i++;
        }
        if (i < nums.length) allNeg = false;

        while (i < nums.length) {
            if (nums[i] >= 0) {
                while (i < nums.length && nums[i] >= 0) {
                    currentMsa += nums[i++];
                }
            } else {
                int negSum = 0;
                while (i < nums.length && nums[i] < 0) {
                    negSum += nums[i++];
                }
                if (maxMsa < currentMsa) {
                    maxMsa = currentMsa;
                }
                if (negSum + currentMsa <= 0) {
                    currentMsa = 0;
                } else {
                    currentMsa += negSum;
                }
            }
        }
        if (maxMsa < currentMsa) {
            maxMsa = currentMsa;
        }

        if (allNeg) return maxNeg;
        else return maxMsa;
    }

    public static void main(String[] args) {
        MaxSubArray msa = new MaxSubArray();

        int[][] tests = {
            {1},
            {1, 2, 3, 4, 5},
            {-2, 1, -3, 4, -1, 2, 1, -5, 4},
            {-2, -1, -3, -4, -6, -7, -5},
            {-2, -1, -3, -4, 1, -6, -7, -5},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(msa.maxSubArray(test));
            System.out.println();
        }
    }
}
