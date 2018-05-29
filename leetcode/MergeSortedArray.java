// https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/587/

class MergeSortedArray {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i1 = m-1;
        int i2 = n-1;
        for (int i = m+n-1; i >=0; i--) {
            if (i2 < 0 || (i1 >= 0 && nums1[i1] > nums2[i2])) {
                nums1[i] = nums1[i1--];
            } else {
                nums1[i] = nums2[i2--];
            }
        }
    }

    public static void main(String[] args) {
        MergeSortedArray msa = new MergeSortedArray();

        int[][][] tests = {
            {{1, 2, 3, 0, 0, 0}, {2, 5, 6}},
            {{0}, {1}},
        };
        for (int[][] test : tests) {
            System.out.println(test[0].length-test[1].length + " " + test[1].length);
            for (int i = 0; i < test[0].length - test[1].length; i++) {
                System.out.print(test[0][i] + " ");
            }
            System.out.println();
            for (int i = 0; i < test[1].length; i++) {
                System.out.print(test[1][i] + " ");
            }
            System.out.println();
            msa.merge(test[0], test[0].length-test[1].length, test[1], test[1].length);
            for (int i = 0; i < test[0].length; i++) {
                System.out.print(test[0][i] + " ");
            }
            System.out.println();
        }
    }
}
