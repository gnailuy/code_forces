// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/770/

class RotateImage {
    public void rotate(int[][] matrix) {
        if (matrix.length <= 1) {
            return;
        }
        assert matrix.length == matrix[0].length;

        for (int l = 0; l < matrix.length/2; l++) { // The layer
            for (int j = l; j < matrix.length - l - 1; j++) { // The point x axis
                int top = matrix[l][j];
                int left = matrix[j][matrix.length-l-1];
                int down = matrix[matrix.length-l-1][matrix.length-j-1];
                int right = matrix[matrix.length-j-1][l];
                matrix[l][j] = right;
                matrix[j][matrix.length-l-1] = top;
                matrix[matrix.length-l-1][matrix.length-j-1] = left;
                matrix[matrix.length-j-1][l] = down;
            }
        }
    }

    public static void main(String[] args) {
        RotateImage ri = new RotateImage();

        int[][][] tests = {
            {
                {1, 2, 3},
                {4, 5, 6},
                {7, 8, 9}
            },
            {
                { 5,  1,  9, 11},
                { 2,  4,  8, 10},
                {13,  3,  6,  7},
                {15, 14, 12, 16}
            },
            {
                { 1,  2,  3,  4,  5},
                { 6,  7,  8,  9, 10},
                {11, 12, 13, 14, 15},
                {16, 17, 18, 19, 20},
                {21, 22, 23, 24, 25},
            }
        };
        for (int[][] test : tests) {
            System.out.println(test.length + "x" + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            ri.rotate(test);
            System.out.println(test.length + "x" + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            System.out.println();
        }
    }
}
