// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/804/

class SearchRotatedSortedArray {
    public int findStart(int[] nums, int left, int right) {
        if (left == right) {
            if (left > 0 && nums[left-1] > nums[left]) return left;
            else if (left < nums.length-1 && nums[left] > nums[left+1]) return left+1;
            else return left;
        } else if (right - left == 1) {
            if (nums[left] > nums[right]) return right;
            else if (right == nums.length-1) return 0;
            else return left;
        } else {
            int middle = (left+right)/2;
            if (nums[middle] > nums[0]) {
                return findStart(nums, middle, right);
            } else { // nums[middle] <= nums[0]
                return findStart(nums, left, middle);
            }
        }
    }

    public int binarySearch(int[] nums, int left, int right, int target) {
        while (left <= right) {
            int middle = (left+right)/2;
            if (nums[middle] == target) {
                return middle;
            } else if (nums[middle] < target) {
                left = middle+1;
            } else { // nums[middle] > target
                right = middle-1;
            }
        }

        return -1;
    }

    public int search(int[] nums, int target) {
        if (nums.length <= 0) return -1;
        int start = findStart(nums, 0, nums.length-1);
        if (start > 0 && nums[0] <= target && target <= nums[start-1]) {
            // Search 0 to start-1
            return binarySearch(nums, 0, start-1, target);
        } else {
            // Search start to end
            return binarySearch(nums, start, nums.length-1, target);
        }
    }

    public static void main(String[] args) {
        SearchRotatedSortedArray srsa = new SearchRotatedSortedArray();

        int[][][] tests = {
            {{4, 5, 6, 7, 0, 1, 2}, {0}},
            {{4, 5, 6, 7, 0, 1, 2}, {3}},
            {{0, 1, 2, 4, 5, 6, 7}, {2}},
            {{0, 1, 2, 4, 5, 6, 7}, {3}},
            {{3, 1}, {1}},
            {{3, 5, 1}, {1}},
            {{5, 1, 3}, {5}},
            {{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, {25}},
        };
        for (int[][] test : tests) {
            int[] nums = test[0];
            int target = test[1][0];

            System.out.println(nums.length + " " + target);
            for (int i = 0; i < nums.length; i++) {
                System.out.print(nums[i] + " ");
            }
            System.out.println();

            System.out.println(srsa.search(nums, target));
            System.out.println();
        }
    }
}
