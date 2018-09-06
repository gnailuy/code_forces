// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/832/

class FirstMissingPositive {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;
        for (int i = 0; i < n; i++) {
            if (nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i]) {
                int origin = nums[nums[i]-1];
                nums[nums[i]-1] = nums[i];
                while (origin > 0 && origin <= n && nums[origin-1] != origin) {

                    int tmp = origin;
                    origin = nums[origin-1];
                    nums[tmp-1] = tmp;
                }
            }
        }
        for (int i = 1; i <= n; i++) {
            if (nums[i-1] != i) {
                return i;
            }
        }
        return n+1;
    }

    public static void main(String[] args) {
        FirstMissingPositive fmp = new FirstMissingPositive();

        int[][] tests = {
            {1},
            {1, 2, 0},
            {3, 4, -1, 1},
            {7, 8, 9, 11, 12},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println(fmp.firstMissingPositive(test));
            System.out.println();
        }
    }
}
