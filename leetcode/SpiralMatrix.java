// https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/828/

import java.util.LinkedList;
import java.util.List;

class SpiralMatrix {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new LinkedList<Integer>();

        int m = matrix.length;
        if (m > 0) {
            int n = matrix[0].length;
            if (n > 0) {
                int l = 0;
                for (l = 0; l < m/2 && l < n/2; l++) {
                    for (int j = l; j < n-l-1; j++) {
                        result.add(matrix[l][j]);
                    }
                    for (int i = l; i < m-l-1; i++) {
                        result.add(matrix[i][n-l-1]);
                    }
                    for (int j = n-l-1; j > l; j--) {
                        result.add(matrix[m-l-1][j]);
                    }
                    for (int i = m-l-1; i > l; i--) {
                        result.add(matrix[i][l]);
                    }
                }
                if (m > n && n%2 != 0) {
                    for (int i = l; i < m-l; i++) {
                        result.add(matrix[i][l]);
                    }
                } else if (m%2 != 0) {
                    for (int j = l; j < n-l; j++) {
                        result.add(matrix[l][j]);
                    }
                }
            }
        }

        return result;
    }

    public static void main(String[] args) {
        SpiralMatrix sm = new SpiralMatrix();

        int[][][] tests = {
            {
                {1, 2, 3},
                {4, 5, 6},
                {7, 8, 9},
            },
            {
                {1,  2,  3,  4},
                {5,  6,  7,  8},
                {9, 10, 11, 12},
            },
            {
                { 1,  2,  3},
                { 4,  5,  6},
                { 7,  8,  9},
                {10, 11, 12},
            },
            {
                { 5,  1,  9, 11},
                { 2,  4,  8, 10},
                {13,  3,  6,  7},
                {15, 14, 12, 16},
            },
            {
                { 1,  2,  3,  4,  5},
                { 6,  7,  8,  9, 10},
                {11, 12, 13, 14, 15},
                {16, 17, 18, 19, 20},
                {21, 22, 23, 24, 25},
            },
            {
                { 1,  2},
                { 6,  7},
                {11, 12},
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
            List<Integer> result = sm.spiralOrder(test);
            System.out.println(result.size());
            for (Integer i: result) {
                System.out.print(i + " ");
            }
            System.out.println();
        }
    }
}
