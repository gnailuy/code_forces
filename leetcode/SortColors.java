// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/798/

class SortColors {
    public void sortColors(int[] nums) {
        int n = nums.length;
        int i = 0; // Next of last 0
        int j = 0; // Next of last 1
        int k = n-1; // Previous of first 2

        while (i < n && nums[i] == 0) i++;
        j = i;
        while (j < n && nums[j] == 1) j++;
        while (k >= 0 && nums[k] == 2) k--;

        while (j <= k) {
            if (nums[j] == 0) {
                if (i < j) {
                    swap(nums, i, j);
                    i++;
                } else { // i == j
                    i++;
                    j++;
                }
            } else if (nums[j] == 1) {
                j++;
            } else if (nums[j] == 2) {
                swap(nums, j, k);
                k--;
            }
        }
    }

    public void swap(int[] nums, int m, int n) {
        int tmp = nums[m];
        nums[m] = nums[n];
        nums[n] = tmp;
    }

    public static void main(String[] args) {
        SortColors sc = new SortColors();

        int[][] tests = {
            {2, 0, 2, 1, 1, 0},
            {2, 0, 1},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            sc.sortColors(test);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
