#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    char black_hole;
    char ** board = (char **) malloc (sizeof(char *) * n);
    for (int i = 0; i < n; i++) {
        board[i] = (char *) malloc (sizeof(char) * n);
        for (int j = 0; j < n; j++) {
            scanf ("%c", &board[i][j]);
        }
        scanf ("%c", &black_hole); // The tailing '\n'
    }

    bool ok = true;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            int os = 0;
            if (i > 0 && 'o' == board[i - 1][j]) {
                os++;
            }
            if (i < n - 1 && 'o' == board[i + 1][j]) {
                os++;
            }
            if (j > 0 && 'o' == board[i][j - 1]) {
                os++;
            }
            if (j < n - 1 && 'o' == board[i][j + 1]) {
                os++;
            }
            if (os % 2 != 0) {
                ok = false;
                break;
            }
        }
    }

    if (ok) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    for (int i = 0; i < n; i++) {
        free (board[i]);
    }
    free (board);
    return 0;
}

