// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/

class MoveZeros {
    public void moveZeros(int[] nums) {
        int i = 0;
        while (i < nums.length && nums[i] != 0) i++;
        int j = i + 1;
        while (j < nums.length) {
            if (0 != nums[j]) {
                nums[i++] = nums[j];
            }
            j++;
        }
        while (i < nums.length) {
            nums[i++] = 0;
        }
    }

    public static void main(String[] args) {
        MoveZeros mz = new MoveZeros();

        int[][] tests = {
            {0, 1, 0, 3, 12},
            {1, 2, 3, 4, 5},
            {1, 0, 0, 0, 12},
            {1, 0, 2, 0, 12},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            mz.moveZeros(test);
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
        }
    }
}
