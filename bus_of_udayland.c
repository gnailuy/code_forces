#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define LINE_LENGTH 5

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    bool found_seats = false;
    char * row = (char *) malloc (sizeof(char) * (LINE_LENGTH + 2));
    char ** seats = (char **) malloc (sizeof(char *) * n);
    for (int i = 0; i < n; i++) {
        fgets(row, LINE_LENGTH + 2, stdin); // '\n'
        seats[i] = (char *) malloc (sizeof(char) * LINE_LENGTH + 1);

        if (found_seats) {
            for (int j = 0; j < LINE_LENGTH; j++) {
                seats[i][j] = row[j];
            }
        } else if ('O' == row[0] && 'O' == row[1]) {
            seats[i][0] = seats[i][1] = '+';
            for (int j = 2; j < LINE_LENGTH; j++) {
                seats[i][j] = row[j];
            }
            found_seats = true;
        } else if ('O' == row[3] && 'O' == row[4]) {
            seats[i][3] = seats[i][4] = '+';
            for (int j = 0; j < LINE_LENGTH - 2; j++) {
                seats[i][j] = row[j];
            }
            found_seats = true;
        } else {
            for (int j = 0; j < LINE_LENGTH; j++) {
                seats[i][j] = row[j];
            }
        }
        seats[i][LINE_LENGTH] = '\0';
    }

    if (found_seats) {
        printf ("YES\n");
        for (int i = 0; i < n; i++) {
            printf ("%s\n", seats[i]);
        }
    } else {
        printf ("NO\n");
    }

    for (int i = 0; i < n; i++) {
        free (seats[i]);
    }
    free (seats);
    free (row);
    return 0;
}

