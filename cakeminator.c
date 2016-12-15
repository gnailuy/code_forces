#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int r, c; char in;
    scanf ("%d %d\n", &r, &c);

    char ** cake = (char **) malloc (sizeof(int *) * r);
    for (int i = 0; i < r; i++) {
        cake[i] = (char *) malloc (sizeof(int) * c);
        for (int j = 0; j < c; j++) {
            scanf ("%c", &in);
            cake[i][j] = in;
        }
        scanf ("%c", &in);
    }

    for (int i = 0; i < r; i++) {
        bool empty = true;
        for (int j = 0; j < c; j++) {
            if ('S' == cake[i][j]) {
                empty = false;
            }
        }
        if (empty) {
            for (int j = 0; j < c; j++) {
                cake[i][j] = ' ';
            }
        }
    }
    for (int i = 0; i < c; i++) {
        bool empty = true;
        for (int j = 0; j < r; j++) {
            if ('S' == cake[j][i]) {
                empty = false;
            }
        }
        if (empty) {
            for (int j = 0; j < r; j++) {
                cake[j][i] = ' ';
            }
        }
    }

    int count = 0;
    for (int i = 0; i < r; i++) {
        for (int j = 0; j < c; j++) {
            if (' ' == cake[i][j]) {
                count++;
            }
        }
    }

    printf ("%d\n", count);

    for (int i = 0; i < r; i++) {
        free (cake[i]);
    }
    free (cake);
    return 0;
}

