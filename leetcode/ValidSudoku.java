// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/

import java.util.HashSet;

class ValidSudoku {
    public boolean isValidSudoku(char[][] board) {
        assert 9 == board.length;
        assert 9 == board[0].length;

        HashSet<Character> row = new HashSet<Character>();
        HashSet<Character> column = new HashSet<Character>();
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                // Row
                if ('.' != board[i][j]) {
                    if (row.contains(board[i][j])) {
                        return false;
                    } else {
                        row.add(board[i][j]);
                    }
                }
                // Column
                if ('.' != board[j][i]) {
                    if (column.contains(board[j][i])) {
                        return false;
                    } else {
                        column.add(board[j][i]);
                    }
                }
            }
            row.clear();
            column.clear();
        }

        // Cells
        HashSet<Character> cell = row; // Reuse empty hash set
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                for (int r = i*3; r < i*3+3; r++) {
                    for (int c = j*3; c < j*3+3; c++) {
                        if ('.' != board[r][c]) {
                            if (cell.contains(board[r][c])) {
                                return false;
                            } else {
                                cell.add(board[r][c]);
                            }
                        }
                    }
                }
                cell.clear();
            }
        }

        return true;
    }

    public static void main(String[] args) {
        ValidSudoku vs = new ValidSudoku();

        char[][][] tests = {
            {
                {'5','3','.','.','7','.','.','.','.'},
                {'6','.','.','1','9','5','.','.','.'},
                {'.','9','8','.','.','.','.','6','.'},
                {'8','.','.','.','6','.','.','.','3'},
                {'4','.','.','8','.','3','.','.','1'},
                {'7','.','.','.','2','.','.','.','6'},
                {'.','6','.','.','.','.','2','8','.'},
                {'.','.','.','4','1','9','.','.','5'},
                {'.','.','.','.','8','.','.','7','9'}
            },
            {
                {'8','3','.','.','7','.','.','.','.'},
                {'6','.','.','1','9','5','.','.','.'},
                {'.','9','8','.','.','.','.','6','.'},
                {'8','.','.','.','6','.','.','.','3'},
                {'4','.','.','8','.','3','.','.','1'},
                {'7','.','.','.','2','.','.','.','6'},
                {'.','6','.','.','.','.','2','8','.'},
                {'.','.','.','4','1','9','.','.','5'},
                {'.','.','.','.','8','.','.','7','9'}
            },
            {
                {'5','3','.','.','7','.','.','.','.'},
                {'6','.','.','1','9','5','.','.','.'},
                {'.','9','8','7','.','.','.','6','.'},
                {'8','.','.','.','6','.','.','.','3'},
                {'4','.','.','8','.','3','.','.','1'},
                {'7','.','.','.','2','.','.','.','6'},
                {'.','6','.','.','.','.','2','8','.'},
                {'.','.','.','4','1','9','.','.','5'},
                {'.','.','.','.','8','.','.','7','9'}
            }
        };
        for (char[][] test : tests) {
            System.out.println(test.length + "x" + test[0].length);
            for (int i = 0; i < test.length; i++) {
                for (int j = 0; j < test[i].length; j++) {
                    System.out.print(test[i][j] + " ");
                }
                System.out.println();
            }
            System.out.println(vs.isValidSudoku(test));
            System.out.println();
        }
    }
}
