// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/

class RemoveDuplicate {
    public int removeDuplicates(int[] nums) {
        if (nums.length <= 1) {
            return nums.length;
        }

        int tailIndex = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[tailIndex-1]) {
                if (i == tailIndex) {
                    tailIndex++;
                } else {
                    nums[tailIndex++] = nums[i];
                }
            }
        }
        return tailIndex;
    }

    public static void main(String[] args) {
        RemoveDuplicate rd = new RemoveDuplicate();

        int[][] tests = {
            {1, 1, 2},
            {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(rd.removeDuplicates(test));
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
