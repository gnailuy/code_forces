// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/

class SetMatrixZeroes {
    public void setZeroes(int[][] matrix) {
        int n = matrix.length;
        if (n <= 0) {
            return;
        }
        int m = matrix[0].length;
        if (m <= 0) {
            return;
        }

        boolean[] columns = new boolean[m];
        for (int j = 0; j < m; j++) {
            columns[j] = false;
        }
        for (int i = 0; i < n; i++) {
            boolean hasZero = false;
            for (int j = 0; j < m; j++) {
                if (0 == matrix[i][j]) {
                    hasZero = true;
                    columns[j] = true;
                }
            }
            if (hasZero) {
                for (int j = 0; j < m; j++) {
                    matrix[i][j] = 0;
                }
            }
        }
        for (int j = 0; j < m; j++) {
            if (columns[j]) {
                for (int i = 0; i < n; i++) {
                    matrix[i][j] = 0;
                }
            }
        }
    }

    public static void main(String[] args) {
        SetMatrixZeroes smz = new SetMatrixZeroes();

        int[][][] tests = {
            {
                {0, 1},
            },
            {
                {0},
                {1},
            },
            {
                {1, 1, 1},
                {1, 0, 1},
                {1, 1, 1}
            },
            {
                {0, 1, 2, 0},
                {3, 4, 5, 2},
                {1, 3, 1, 5},
            }
        };
        for (int[][] test : tests) {
            System.out.println(test.length + " " + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            System.out.println();
            smz.setZeroes(test);
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
