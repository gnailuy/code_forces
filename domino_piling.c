#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void swap(int *m, int *n) {
    *m = *m ^ *n;
    *n = *m ^ *n;
    *m = *m ^ *n;
}

void print_board(int **board, int m, int n) {
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            printf("%d ", board[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

int **generate_board(int m, int n) {
    int **board = (int **) malloc(sizeof(int*) * m);
    for (int i = 0; i < m; i++) {
        board[i] = (int *) malloc(sizeof(int) * n);
        memset(board[i], 0, sizeof(int) * n);
    }
    return board;
}

void destroy_board(int **board, int m, int n) {
    for (int i = 0; i < m; i++) {
        free (board[i]);
    }
    free (board);
}

bool greedily_put_one(int **board, int m, int n) {
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            if (1 == board[i][j]) {
                continue;
            } else if (j < n - 1) {
                board[i][j] = 1;
                board[i][j + 1] = 1;
                return true;
            } else if (i < m - 1) {
                board[i][j] = 1;
                board[i + 1][j] = 1;
                return true;
            } else {
                return false;
            }
        }
    }
    return false;
}

int main(int argc, char * argv[]) {
    int m, n;
    scanf("%d %d", &m, &n);

    if (m > n) {
        swap(&m, &n);
    }
    int **board = generate_board(m, n);
    // print_board(board, m, n);

    int count = 0;
    bool posso_put;
    do {
        posso_put = greedily_put_one(board, m, n);
        if (posso_put) {
            // print_board(board, m, n);
            count++;
        }
    } while (posso_put);

    printf("%d\n", count);
    
    destroy_board (board, m, n);
    return 0;
}

