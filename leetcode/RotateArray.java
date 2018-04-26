class RotateArray {
    public void rotate(int[] nums, int k) {
        int[] extra = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            extra[i] = nums[i];
        }
        for (int i = 0; i < nums.length; i++) {
            nums[(i+k)%nums.length] = extra[i];
        }
    }

    public static void main(String[] args) {
        RotateArray ra = new RotateArray();

        int[][][] tests = {
            {{}, {1}},
            {{1}, {1}},
            {{1, 2}, {5}},
            {{1, 2, 3}, {0}},
            {{1, 2, 3, 4, 5, 6, 7}, {3}}
        };
        for (int[][] test : tests) {
            int[] nums = test[0];
            int k = test[1][0];

            System.out.println(nums.length);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();
            ra.rotate(nums, k);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
