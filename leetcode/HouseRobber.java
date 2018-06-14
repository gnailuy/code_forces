// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/576/

class HouseRobber {
    private int max(int a, int b) {
        if (a >= b) return a;
        else return b;
    }

    public int rob(int[] nums) {
        int n = nums.length;
        if (0 == n)
            return 0;
        else if (1 == n)
            return nums[0];
        else if (2 == n)
            return max(nums[0], nums[1]);

        int[] max_array = new int[n];
        max_array[0] = nums[0];
        max_array[1] = max(nums[0], nums[1]);

        for (int i = 2; i < n; i++) {
            max_array[i] = max(max_array[i-1], max_array[i-2]+nums[i]);
        }

        return max_array[n-1];
    }

    public static void main(String[] args) {
        HouseRobber hr = new HouseRobber();

        int[][] tests = {
            {1, 2, 3, 1},
            {2, 7, 9, 3, 1},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(hr.rob(test));
            System.out.println();
        }
    }
}
