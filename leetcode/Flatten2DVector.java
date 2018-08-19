// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/806/

class Flatten2DVector {
    int[][] vector = null;
    int m = 0;
    int n = 0;

    public Flatten2DVector(int[][] vector) {
        this.vector = vector;
    }

    public boolean hasNext() {
        while (vector.length > m) {
            if (vector[m].length > n) {
                return true;
            } else { // In case of empty line
                m++;
                n = 0;
            }
        }
        return false;
    }

    public int next() {
        int value = vector[m][n];
        if (vector[m].length-1 == n) {
            m++;
            n = 0;
        } else {
            n++;
        }
        return value;
    }

    public static void main(String[] args) {
        int[][][] tests = {
            {
                {},
                {},
                {1, 2, 3},
                {4, 6},
                {7, 8, 9},
            },
            {
                {},
                {},
                {},
                {1, 4, 7, 11, 15},
                {2, 5, 12, 19},
                {3, 22},
                {},
                {},
                {10, 17, 24},
                {18, 21, 23, 26, 30},
                {},
                {},
            },
        };

        for (int[][] test: tests) {
            System.out.println(test.length);
            for (int i = 0; i < test.length; i++) {
                System.out.print("Line: ");
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            Flatten2DVector fv = new Flatten2DVector(test);

            while (fv.hasNext()) {
                System.out.print(fv.next() + " ");
            }
            System.out.println();
            System.out.println();
        }
    }
}
