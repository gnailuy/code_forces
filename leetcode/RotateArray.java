// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/

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

    public int gcd(int x, int y) {
        int tmp;
        if (x < y) {
            tmp = x;
            x = y;
            y = tmp;
        }
        while (true) {
            if (x % y == 0) {
                return y;
            } else {
                tmp = x % y;
                x = y;
                y = tmp;
            }
        }
    }

    public void rotate_o1(int[] nums, int k) {
        if (nums.length <= 1) {
            return;
        }
        k = k%nums.length;
        if (k == 0) {
            return;
        }

        int unit = gcd(nums.length, k);
        for (int idx = 0; idx < unit; idx++) {
            int extra = nums[idx];
            for (int i = 0; i < nums.length/unit; i++) {
                nums[(idx+k)%nums.length] ^= extra;
                extra ^= nums[(idx+k)%nums.length];
                nums[(idx+k)%nums.length] ^= extra;
                idx = (idx+k)%nums.length;
            }
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

            System.out.println(nums.length + " " + k);
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

        int[][][] tests_o1 = {
            {{1}, {1}},
            {{1, 2, 3, 4, 5, 6}, {2}},
            {{1, 2, 3, 4, 5, 6, 7}, {3}}
        };
        for (int[][] test : tests_o1) {
            int[] nums = test[0];
            int k = test[1][0];

            System.out.println(nums.length + " " + k);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();
            ra.rotate_o1(nums, k);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
