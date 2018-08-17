// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/792/

import java.util.LinkedList;

class NumberOfIslands {
    class Point {
        int x;
        int y;
        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    public int numIslands(char[][] grid) {
        int m = grid.length;
        if (m <= 0) return 0;
        int n = grid[0].length;

        int[][] rep = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                rep[i][j] = grid[i][j] - '0';
            }
        }

        int current = 1;
        LinkedList<Point> stack = new LinkedList<Point>();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (1 == rep[i][j]) {
                    current++;

                    // BFS
                    stack.push(new Point(i, j));
                    while (!stack.isEmpty()) {
                        Point p = stack.pop();
                        rep[p.x][p.y] = current;
                        if (p.x > 0 && 1 == rep[p.x-1][p.y]) stack.push(new Point(p.x-1, p.y));
                        if (p.y > 0 && 1 == rep[p.x][p.y-1]) stack.push(new Point(p.x, p.y-1));
                        if (p.x < m-1 && 1 == rep[p.x+1][p.y]) stack.push(new Point(p.x+1, p.y));
                        if (p.y < n-1 && 1 == rep[p.x][p.y+1]) stack.push(new Point(p.x, p.y+1));
                    }
                }
            }
        }

        return current - 1;
    }

    public static void main(String[] args) {
        NumberOfIslands noi = new NumberOfIslands();

        char[][][] tests = {
            {
                {'1', '1', '1', '1', '0'},
                {'1', '1', '0', '1', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '0', '0', '0'},
            },
            {
                {'1', '1', '0', '0', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '1', '0', '0'},
                {'0', '0', '0', '1', '1'},
            },
        };
        for (char[][] test : tests) {
            System.out.println(test.length + "x" + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            System.out.println(noi.numIslands(test));
        }
    }
}
