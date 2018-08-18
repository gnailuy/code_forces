// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/800/

class KthLargestElementInArray {
    public int partition(int[] nums, int k, int left, int right) {
        int first = nums[left];
        int index = left;
        int i = left+1;
        while (i <= right) {
            if (nums[i] > first) {
                nums[index++] = nums[i];
                nums[i] = nums[index];
                if (index >= i) {
                    i = index+1;
                }
            } else {
                i++;
            }
        }
        nums[index] = first;

        if (index-left == k-1) {
            return nums[index];
        } else if (index-left > k-1) {
            return partition(nums, k, left, index-1);
        } else {
            return partition(nums, k-index+left-1, index+1, right);
        }
    }

    public int findKthLargest(int[] nums, int k) {
        return partition(nums, k, 0, nums.length-1);
    }

    public static void main(String[] args) {
        KthLargestElementInArray klea = new KthLargestElementInArray();

        int[][][] tests = {
            {{3, 2, 1, 5, 6, 4}, {2}},
            {{3, 2, 3, 1, 2, 4, 5, 5, 6}, {4}},
            {{2, 1}, {2}},
            {{7, 6, 5, 4, 3, 2, 1}, {5}},
            {{5, 2, 4, 1, 3, 6, 0}, {4}},
        };
        for (int[][] test : tests) {
            int[] nums = test[0];
            int k = test[1][0];

            System.out.println(nums.length + " " + k);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();

            System.out.println(klea.findKthLargest(nums, k));
            System.out.println();
        }
    }
}
