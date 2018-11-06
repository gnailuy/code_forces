// https://leetcode.com/problems/next-permutation/

class NextPermutation {
    public void nextPermutation(int[] nums) {
        // Find the first i from right that nums[i] < nums[i+1], or -1
        int i = nums.length-1;
        while (i > 0 && nums[i-1] >= nums[i]) {
            i--;
        }
        i -= 1;

        // Replace a[i] with the first one on the right that is larger than a[i]
        int j = nums.length-1;
        if (i >= 0) {
            while (nums[i] >= nums[j]) {
                j--;
            }
            int tmp = nums[i];
            nums[i] = nums[j];
            nums[j] = tmp;
        }

        // Reverse nums[i+1:end]
        i += 1;
        j = nums.length-1;
        while (i < j) {
            int tmp = nums[i];
            nums[i] = nums[j];
            nums[j] = tmp;
            i++;
            j--;
        }
    }

    public static void main(String[] args) {
        NextPermutation np = new NextPermutation();

        int[][] tests = {
            {},
            {1},
            {1, 2},
            {2, 1},
            {1, 2, 3},
            {1, 3, 2},
            {2, 3, 1},
            {1, 5, 1},
        };
        for (int[] test : tests) {
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            np.nextPermutation(test);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}

