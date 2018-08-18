// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/802/

class SearchRange {
    private int searchLeft(int[] nums, int target, int left, int right) {
        if (left >= right) return left;

        int middle = (left+right)/2;
        if (target == nums[middle]) {
            return searchLeft(nums, target, left, middle); // Always have one in range
        } else { // target > nums[middle]
            return searchLeft(nums, target, middle+1, right);
        }
    }

    private int searchRight(int[] nums, int target, int left, int right) {
        if (left >= right) return right;

        int middle = (left+right+1)/2;
        if (target == nums[middle]) {
            return searchRight(nums, target, middle, right); // Always have one in range
        } else { // target < nums[middle]
            return searchRight(nums, target, left, middle-1);
        }
    }

    public int[] searchRange(int[] nums, int target, int left, int right) {
        if (left > right) {
            return new int[]{-1, -1};
        }

        int middle = (left+right)/2;
        if (target == nums[middle]) {
            int leftIndex, rightIndex;
            if (middle <= 0 || nums[middle-1] < target) {
                leftIndex = middle;
            } else {
                leftIndex = searchLeft(nums, target, left, middle-1);
            }
            if (middle >= nums.length-1 || nums[middle+1] > target) {
                rightIndex = middle;
            } else {
                rightIndex = searchRight(nums, target, middle+1, right);
            }
            return new int[]{leftIndex, rightIndex};
        } else if (target < nums[middle]) {
            return searchRange(nums, target, left, middle-1);
        } else { // target > nums[middle]
            return searchRange(nums, target, left+1, right);
        }
    }

    public int[] searchRange(int[] nums, int target) {
        return searchRange(nums, target, 0, nums.length-1);
    }

    public static void main(String[] args) {
        SearchRange sr = new SearchRange();

        int[][][] tests = {
            {{5, 7, 7, 8, 8, 10}, {8}},
            {{5, 7, 7, 8, 8, 10}, {6}},
            {{5, 8, 8, 8, 8, 10}, {8}},
            {{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 8, 10, 10}, {4}},
        };
        for (int[][] test : tests) {
            System.out.println(test[0].length + " " + test[1][0]);
            for (int i = 0; i < test[0].length; i++) {
                System.out.print(test[0][i] + " ");
            }
            System.out.println();

            int[] result = sr.searchRange(test[0], test[1][0]);
            System.out.println(result[0] + " " + result[1]);
            System.out.println();
        }
    }
}
