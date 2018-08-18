// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/806/

class Search2DMatrix {
    public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length;
        if (m <= 0) return false;
        int n = matrix[0].length;
        if (n <= 0) return false;

        int i = 0, j = n-1;
        while (i < m && j >= 0) {
            if (matrix[i][j] == target) {
                return true;
            } else if (matrix[i][j] < target) {
                i++;
            } else { // matrix[i][j] > target
                j--;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Search2DMatrix sm = new Search2DMatrix();

        int[][][] tests = {
            {
                {1, 2, 3},
                {4, 5, 6},
                {7, 8, 9}
            },
            {
                {1,   4,  7, 11, 15},
                {2,   5,  8, 12, 19},
                {3,   6,  9, 16, 22},
                {10, 13, 14, 17, 24},
                {18, 21, 23, 26, 30}
            },
        };
        int[][] queries = {
            {1, 5, 9},
            {5, 20},
        };
        for (int k = 0; k < tests.length; k++) {
            int[][] test = tests[k];
            int[] query = queries[k];
            System.out.println(test.length + "x" + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }

            for (int l = 0; l < query.length; l++) {
                System.out.print(query[l] + ":" + sm.searchMatrix(test, query[l]) + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
