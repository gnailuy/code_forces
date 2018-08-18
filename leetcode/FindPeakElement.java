// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/801/

class FindPeakElement {
    private int findPeakElement(int[] nums, int left, int right) {
        int middle = (left+right)/2;
        if ((middle == 0 || nums[middle-1] < nums[middle])
                && (middle == nums.length-1 || nums[middle+1] < nums[middle])) {
            return middle;
        }

        if (middle > 0 && nums[middle-1] > nums[middle]) {
            return findPeakElement(nums, left, middle-1);
        } else if (middle < nums.length-1 && nums[middle+1] > nums[middle]) {
            return findPeakElement(nums, middle+1, right);
        }
        return -1;
    }

    public int findPeakElement(int[] nums) {
        return findPeakElement(nums, 0, nums.length-1);
    }

    public static void main(String[] args) {
        FindPeakElement fpe = new FindPeakElement();

        int[][] tests = {
            {1},
            {1, 2, 3, 1},
            {1, 2, 1, 3, 5, 6, 4},
            {1, 2, 1, 4, 2, 3},
        };
        for (int[] test : tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print(test[i] + " ");
            }
            System.out.println();

            System.out.println(fpe.findPeakElement(test));
            System.out.println();
        }
    }
}
